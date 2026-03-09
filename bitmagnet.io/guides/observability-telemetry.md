---
title: Observability & Telemetry
description: Observability & Telemetry features in bitmagnet
parent: Guides
layout: default
nav_order: 9
redirect_from:
  - /internals-development/observability-telemetry.html
---

# Observability & Telemetry

## Grafana stack & Prometheus integration

**bitmagnet** can integrate with the [Grafana stack](https://grafana.com/) and [Prometheus](https://prometheus.io/) for monitoring and building observability dashboards for the DHT crawler and other components. The repository no longer ships a ready-made observability stack, but the `/metrics` and `/debug/pprof/*` endpoints are still available for wiring into your own Grafana, Prometheus, Loki, or Pyroscope deployment.

![Grafana dashboard](/assets/images/grafana-1.png)

The example integration includes:

- [Grafana](https://grafana.com/oss/grafana/) - A dashboarding and visualization tool
- [Grafana Agent](https://grafana.com/oss/agent/) - Collects metrics and logs, and forwards them to storage backends
- [Prometheus](https://prometheus.io/) - A time series database for metrics
- [Loki](https://grafana.com/oss/loki/) - A log aggregation system
- [Pyroscope](https://pyroscope.io/) - A continuous profiling tool
- [Postgres exporter](https://github.com/prometheus-community/postgres_exporter) - Exposes Postgres metrics to Prometheus

# Profiling with pprof

**bitmagnet** exposes [Go pprof](https://golang.org/pkg/net/http/pprof/) profiling endpoints at `/debug/pprof/*`, for example:

```sh
go tool pprof http://localhost:3333/debug/pprof/heap
```
