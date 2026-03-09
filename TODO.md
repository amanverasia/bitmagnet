# TODO

## Low Effort, High Value
- [ ] Add a `healthcheck` to the `bitmagnet` service in `docker-compose.yml` using `/status`.
- [ ] Expand `README.md` with local run instructions, Docker usage, and the current fork/module path.
- [ ] Add a sample environment/config file such as `.env.example` or a documented `config.yml` template.
- [ ] Document the purpose of each root Dockerfile (`Dockerfile`, `Dockerfile.scratch`, `Dockerfile_dev`, `ci.Dockerfile`, `goreleaser.Dockerfile`).
- [ ] Fix `AGENTS.md` so it reads cleanly both in the terminal workflow and on GitHub.

## Tooling And Verification
- [ ] Restore automated checks in CI or an equivalent local automation path for `task test`, `task lint`, `task gen`, and Docker smoke tests.
- [ ] Re-enable `lint-golangci` in the default `task lint` flow after fixing the Nix/tooling issue in `Taskfile.yml`.
- [ ] Add an automated smoke test that boots Postgres and verifies the `/status` endpoint.
- [ ] Reintroduce generated-code drift detection after `task gen`.
- [ ] Run and stabilize the full verification path after the fork migration: `go test -v ./...` and `cd webui && CHROME_BIN=chromium npm test`.

## Deployment And Runtime
- [ ] Stop running the default Docker stack as root, or switch bind mounts to named volumes to avoid root-owned local files.
- [ ] Split Docker Compose usage into clearer modes such as `search-only`, `crawler`, and `dev`.
- [ ] Review whether `/metrics` and `/debug/pprof/*` should stay enabled by default in minimal deployments.
- [ ] Decide whether `webui/dist` should stay committed or always be built as part of packaging.
- [ ] Consolidate repeated version/ldflags wiring across the Taskfile and all Dockerfiles.
- [ ] Tighten `.dockerignore` and reduce Docker build context further where possible.
- [ ] Standardize on fewer Docker build paths if some of the current Dockerfiles are redundant.

## Code Structure
- [ ] Make TMDB wiring optional in code, not just by runtime config.
- [ ] Make telemetry wiring optional in code, not just by runtime config.
- [ ] Add explicit tests for `TMDB_ENABLED=false`.
- [ ] Add tests for worker combinations such as `http_server` only, `queue_server` only, and crawler on/off.
- [ ] Implement the configurable panic handler noted in `internal/health/check.go`.
- [ ] Review copied-code areas and decide whether to isolate, replace, upstream, or document them more explicitly.

## Frontend And Dependencies
- [ ] Refresh frontend dependencies and reduce deprecated transitive packages in `webui/package-lock.json`.
- [ ] Review whether the committed built assets in `webui/dist` are still the right workflow for this fork.

## Product And Branding
- [ ] Decide whether the fork should continue using the `bitmagnet` product name.
- [ ] Decide whether links to `bitmagnet.io` should remain or be replaced with fork-owned documentation.
- [ ] Add a clear fork-specific project overview to `README.md` explaining what changed from upstream.

## Notes
- The easiest wins are the docs and Docker healthcheck items at the top of this file.
- The riskiest changes are the code-optionalization work for TMDB/telemetry and any branding or packaging workflow changes.
