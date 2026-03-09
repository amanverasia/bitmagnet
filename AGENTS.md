# Repository Guidelines

## Project Structure & Module Organization
`main.go` is the application entrypoint. The Go module path is `github.com/amanverasia/bitmagnet`. Backend code lives under `internal/` and is grouped by domain (`internal/gql`, `internal/database`, `internal/protocol`, `internal/tmdb`, etc.). Database migrations are in `migrations/`. Shared GraphQL documents and schema inputs live in `graphql/`. The Angular Web UI source is in `webui/src/`; production assets are emitted to `webui/dist/` and embedded into the Go binary. Repository docs and setup guides live in `bitmagnet.io/`.

## Build, Test, and Development Commands
Prefer running tasks through the Nix dev shell to match CI: `nix develop --ignore-environment --command task <name>`.

- `docker compose up --build`: build the `scratch`-style local image and start the minimal `bitmagnet + postgres` stack on `localhost:3333`.
- `task build`: build the Go binary with embedded git version metadata.
- `task test`: run both Go and Web UI test suites.
- `task lint`: run Web UI linting and repository-wide Prettier checks.
- `task gen`: regenerate Go, GraphQL, protobuf, mocks, classifier schema, and Web UI GraphQL code.
- `task migrate`: apply local database migrations.
- `task install-webui` then `task serve-webui`: install frontend dependencies and start the Angular dev server.

## Coding Style & Naming Conventions
Follow `.editorconfig`: UTF-8, LF endings, spaces, and 2-space indentation. Keep Go packages lowercase and organized by domain; test files use `*_test.go`. Go formatting is enforced by `gofmt`, `gofumpt`, `goimports`, `gci`, and `golines`, with `golangci-lint` enforcing additional rules. In `webui/`, use Angular’s established naming: kebab-case files and suffixes such as `*.component.ts`, `*.service.ts`, and `*.spec.ts`. Use Prettier and ESLint before opening a PR.

## Testing Guidelines
Run `go test -v ./...` for backend coverage and `cd webui && CHROME_BIN=chromium npm test` for Angular unit tests. Add or update tests for every behavior change: Go tests should sit beside the package they cover, and frontend tests should mirror the component/service filename with `.spec.ts`. There is no published coverage threshold, but CI expects touched areas to remain covered and generated artifacts to stay committed.

## Configuration Tips
The checked-in [docker-compose.yml](/home/amanverasia/Projects/bitmagnet/docker-compose.yml) is intentionally minimal and disables TMDB by default with `TMDB_ENABLED=false`. Persist config under `./config` and Postgres data under `./data/postgres`. If you change embedded frontend or generated GraphQL code, rebuild the image before testing the container flow.

## Work Tracking
Keep [TODO.md](/home/amanverasia/Projects/bitmagnet/TODO.md) updated for any multi-step or repo-wide change. Use it to track migration phases, verification status, and any follow-up work that should survive across sessions.

## Commit & Pull Request Guidelines
Recent history favors short, imperative commit subjects, with optional scoped prefixes when useful, for example `fix(search): query parser bug (#423)` or `feat(webui): add catalan i18n (#404)`. Keep commits focused and include regenerated files when `task gen` or translation extraction changes output. PRs should describe the behavior change, link the related issue, note migrations or generated-code updates, and include screenshots for Web UI changes. Redact info hashes and content metadata from logs or screenshots before sharing them.
