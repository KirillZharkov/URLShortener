# Copilot instructions for SimpleRestAPIURLShortener

This repository is a small Go REST API that shortens URLs. Use these concise rules when editing or extending the codebase.

- Project layout
  - Entry point: `cmd/url_shortener/main.go` — the binary initialization follows this sequence: config -> logger -> storage -> router -> server. Follow this order for new features that touch initialization.
  - Single-binary pattern: put executable code under `cmd/url_shortener` and library code under top-level packages (e.g., `internal`, `pkg`) if added.

- What to change and how
  - When adding services (storage, auth, metrics), add package-level constructors like `NewStorage(...)` and wire them in `cmd/url_shortener/main.go` in the same initialization order.
  - Keep `main.go` minimal: it should only orchestrate initialization and start the HTTP server; move business logic to packages.

- Build & run (Windows PowerShell)
  - Run without module support (if no `go.mod`) or with module support if present.
  - Quick run (recommended during development):

```powershell
# run directly
go run ./cmd/url_shortener
```

  - Build a binary:

```powershell
# produce binary in current folder
go build -o bin/url_shortener ./cmd/url_shortener
```

  - If you see module errors, initialize a module at repo root and tidy:

```powershell
# only if go.mod is missing
go mod init github.com/<your-username>/SimpleRestAPIURLShortener
go mod tidy
```

- Tests and linting
  - No tests are present in the repo root currently; add package tests near the package under test and run:

```powershell
go test ./...
```

  - Use `gofmt`/`go vet` for formatting and basic checks. Add `golangci-lint` config if introducing linting.

- Project-specific patterns discovered
  - Initialization order is important and documented in `cmd/url_shortener/main.go` comments — preserve that sequence when adding new global services.
  - No `go.mod` was detected at the time of analysis; assume the project may be a plain GOPATH-style layout. If you convert to modules, update CI/README accordingly.

- Integration & external deps
  - Storage and logger are referenced in `main.go` comments but not implemented; when adding storage, clearly document connection strings in a config package and expose a `Close()` or `Shutdown()` method for graceful termination.
  - Keep external network or DB credentials out of commits; prefer environment variables and a `config` package to load them.

- Small examples
  - Wiring example (pseudocode to follow the repo pattern):

```go
// in main.go
cfg := config.Load()
logger := logger.New(cfg.Logger)
storage := storage.New(cfg.Storage, logger)
router := api.NewRouter(storage, logger)
server := http.NewServer(cfg.Server, router)
server.Start()
```

- What not to change
  - Don't fold complex business logic into `main.go`. Move it into packages and keep `main.go` as the orchestrator.

- When uncertain, ask the user
  - If the module path, preferred storage backend (in-memory / sqlite / postgres), or CI workflow are not obvious, request them before making schema-level changes.

If any section is unclear or you want examples for adding a `storage` package, let me know and I will update this file with concrete code snippets and tests.
