package healthcheckcmd

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/amanverasia/bitmagnet/internal/httpserver"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	HTTPServerConfig httpserver.Config
}

type Result struct {
	fx.Out
	Command *cli.Command `group:"commands"`
}

func New(p Params) Result {
	cmd := &cli.Command{
		Name:  "healthcheck",
		Usage: "Checks whether the local HTTP server is responding on /status",
		Flags: []cli.Flag{
			&cli.DurationFlag{
				Name:  "timeout",
				Usage: "HTTP probe timeout",
				Value: 5 * time.Second,
			},
		},
		Action: func(ctx *cli.Context) error {
			timeout := ctx.Duration("timeout")
			url := statusURL(p.HTTPServerConfig.LocalAddress)
			client := &http.Client{Timeout: timeout}

			req, err := http.NewRequestWithContext(ctx.Context, http.MethodGet, url, nil)
			if err != nil {
				return err
			}

			res, err := client.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return fmt.Errorf("healthcheck failed: %s", res.Status)
			}

			return nil
		},
	}

	return Result{Command: cmd}
}

func statusURL(localAddress string) string {
	host, port, err := net.SplitHostPort(localAddress)
	if err != nil {
		return "http://127.0.0.1:3333/status"
	}

	if host == "" || host == ":" || host == "::" || host == "0.0.0.0" {
		host = "127.0.0.1"
	}

	if strings.Contains(host, ":") && !strings.HasPrefix(host, "[") {
		host = "[" + host + "]"
	}

	return fmt.Sprintf("http://%s:%s/status", host, port)
}
