# TODO

## Fork Independence

### Phase 1: Human-facing references
- [x] Replace repository links in docs and site metadata with `https://github.com/amanverasia/bitmagnet`.
- [x] Remove published image references that still point to the old registry namespace.
- [x] Replace remote install guidance that depends on the upstream module path with local checkout guidance.
- [x] Update contributor guidance to note that this file should be kept current during the fork migration.

### Phase 2: Module identity
- [x] Change the Go module path to `github.com/amanverasia/bitmagnet`.
- [x] Rewrite internal Go imports and build ldflags that still reference the upstream module path.
- [x] Rebuild generated code and verify `docker compose up --build` plus key app endpoints.

## Current Notes
- Phase 1 is documentation and metadata only.
- Phase 2 is a repo-wide code migration and should be done in one pass to avoid a half-renamed module state.
