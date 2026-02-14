# go-nuts

Starter template for Go API projects with a clean, scalable structure based on the Go modules layout guidance for server projects (`cmd/` + `internal/`).

This repository is intentionally organized so new features can be added without turning the codebase into a monolith of handlers.

## Goals of this structure

- Keep binaries (entrypoints) separate from business logic.
- Keep app logic private to this module using `internal/`.
- Make it obvious where to add new code (handlers, services, repositories, models).
- Keep transport concerns (`http`) separated from domain/use-case logic.

## Current project structure

```text
.
├── go.mod
├── go.sum
├── cmd/
│   ├── api-server/
│   │   └── main.go
│   └── metrics-analyzer/
│       └── main.go
└── internal/
    ├── api/
    │   ├── handler/
    │   │   ├── health.go
    │   │   └── orders.go
    │   └── router/
    │       └── router.go
    ├── auth/
    │   └── middleware.go
    ├── config/
    │   └── config.go
    ├── logger/
    │   └── logger.go
    ├── metrics/
    │   └── analyzer.go
    ├── model/
    │   └── response.go
    ├── orders/
    │   ├── order.go
    │   ├── repository.go
    │   └── service.go
    └── server/
        └── server.go
```

## How the app works (high level)

1. `cmd/api-server/main.go` boots the app:
   - loads configuration,
   - initializes `zap` logger,
   - starts HTTP server,
   - performs graceful shutdown on SIGINT/SIGTERM.

2. `internal/server/server.go` owns server lifecycle.

3. `internal/api/router/router.go` wires routes and middleware.

4. `internal/api/handler/*` handles HTTP transport only:
   - parse request,
   - call service,
   - return response.

5. Feature logic lives in feature packages (example: `internal/orders`):
   - domain types (`order.go`),
   - repository interface + adapters (`repository.go`),
   - use-cases/services (`service.go`).

## Where to implement new things

Use this rule of thumb:

- **New endpoint** → `internal/api/handler/` + register in `internal/api/router/router.go`
- **New business rule/use-case** → `internal/<feature>/service.go`
- **New data source (DB, external API, cache)** → adapter implementing `internal/<feature>/repository.go` interface
- **New shared HTTP middleware** → `internal/auth/` or new `internal/middleware/` package
- **New app-wide config/env var** → `internal/config/config.go`
- **New command/binary** (worker, migration runner, batch job) → `cmd/<new-command>/main.go`

## How to implement a new feature (recommended flow)

Example: add a `customers` feature.

1. Create `internal/customers/` with:
   - `customer.go` (domain model)
   - `repository.go` (interface + initial adapter)
   - `service.go` (use-cases)

2. Add HTTP handler in `internal/api/handler/customers.go`.

3. Wire dependencies and route in `internal/api/router/router.go`.

4. Keep handler thin:
   - no DB calls directly,
   - no business rules directly.

5. Add tests:
   - service tests in `internal/customers/*_test.go`
   - handler/router tests in `internal/api/.../*_test.go`

## Existing API endpoints

- `GET /healthz` (public)
- `GET /v1/ping` (requires `X-API-Key`)
- `GET /v1/orders` (requires `X-API-Key`)

### Quick test with curl

```bash
curl -i http://localhost:8080/healthz
curl -i -H 'X-API-Key: local-dev' http://localhost:8080/v1/ping
curl -i -H 'X-API-Key: local-dev' http://localhost:8080/v1/orders
```

## Run commands

Run API server:

```bash
go run ./cmd/api-server
```

Run metrics analyzer:

```bash
go run ./cmd/metrics-analyzer ./README.md
```

## Makefile shortcuts

Common local commands:

```bash
make help
make tidy
make fmt
make vet
make test
make run
make build
```

Build artifact from `make build`:

- `bin/api-server`

## Docker (basic)

Build image:

```bash
make docker-build
```

Run container:

```bash
make docker-run
```

Or with raw Docker commands:

```bash
docker build -t go-nuts:latest .
docker run --rm -p 8080:8080 --name go-nuts go-nuts:latest
```

## Configuration

Environment variables used by API server:

- `APP_ENV` (default: `development`)
- `HTTP_ADDR` (default: `:8080`)
- `HTTP_SHUTDOWN_TIMEOUT` (default: `10s`)
- `HTTP_READ_HEADER_TIMEOUT` (default: `5s`)

Example:

```bash
APP_ENV=production HTTP_ADDR=:9090 go run ./cmd/api-server
```

## Logging

- Logger is initialized in `internal/logger/logger.go`.
- Uses `zap`:
  - development config when `APP_ENV=development`
  - production config otherwise

## Best practices used in this template

- Clear separation of responsibilities by package.
- `internal/` to prevent accidental external coupling.
- Interface boundary at repository layer.
- Graceful server shutdown.
- Structured logging with `zap`.
- Multiple binaries pattern through `cmd/`.

## When to split into another module

If some package becomes generally reusable across multiple repositories, extract that package to a separate Go module.

Keep this repository focused on server application concerns.