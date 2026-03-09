# Repository Guidelines

## Project Structure & Module Organization
`main.go` is the application entrypoint. Backend code lives under `internal/` and is grouped by domain (`internal/gql`, `internal/database`, `internal/protocol`, `internal/tmdb`, etc.). Database migrations are in `migrations/`. Shared GraphQL documents and schema inputs live in `graphql/`. The Angular Web UI is in `webui/src/`; built assets go to `webui/dist/`. Supporting docs and the marketing/site content live in `bitmagnet.io/`.

## Build, Test, and Development Commands
Prefer running tasks through the Nix dev shell to match CI: `nix develop --ignore-environment --command task <name>`.

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

## Commit & Pull Request Guidelines
Recent history favors short, imperative commit subjects, with optional scoped prefixes when useful, for example `fix(search): query parser bug (#423)` or `feat(webui): add catalan i18n (#404)`. Keep commits focused and include regenerated files when `task gen` or translation extraction changes output. PRs should describe the behavior change, link the related issue, note migrations or generated-code updates, and include screenshots for Web UI changes. Redact info hashes and content metadata from logs or screenshots before sharing them.
