## Build & Run
Succinct rules for how to BUILD the project:
- Build binary: `go build -o gocv .`
- Run CLI mode: `./gocv` (Reads `./content`, writes to `./output`, exits)
- Run Serve mode: `./gocv serve` (Starts HTTP server, watches `./content`, live updates)
- Configuration: Edit `config.yaml` for template name and HTTP port. Do not use CLI flags.

## Validation
Run these after implementing to get immediate feedback:
- Tests: `go test ./... -v`
- Typecheck: `go vet ./...`
- Lint: `golangci-lint run` (if available)
- Format: `go fmt ./...`

## Operational Notes
Succinct learnings about how to RUN the project:
- The binary must be standalone (static linking preferred).
- Do not introduce external dependencies for PDF generation (no wkhtmltopdf, no chrome).
- Graceful shutdown is required for `serve` mode (handle SIGINT/SIGTERM).
- Assume reverse proxy handles SSL; server runs on HTTP only.
- Update `PLAN.md` at the end of every session.

### Codebase Patterns
- Use `html/template` for rendering.
- Use `goldmark` or similar for Markdown parsing.
- Use `fsnotify` or similar for file watching in `serve` mode.
- Keep `main.go` clean; delegate logic to packages (`pkg/` or `internal/`).
- Config loading should happen at startup; validate required fields.
- Error handling should be explicit; avoid panics in production paths.
