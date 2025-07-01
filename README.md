# init-microservice-go

I'm mostly from a Java background, so you can find most of my development/design thinking
in [init-microservice](https://github.com/harmelodic/init-microservice).

However, a bunch of the world now is built in other languages. For me, Go and Rust are the other languages that I would
consider using to develop services.

This repo represents a frame for how I'd go about building services in Go.

## Showcases

Project configuration:

- [x] README
- [x] .gitignore
- [x] Auto-updates (Renovate)
- [x] Dependency Management (go.mod)
- [x] Packaged according to [golang-standard's project-layout](https://github.com/golang-standards/project-layout)

Application configuration:

- [x] Dependency Injection (Plain old root DI, could have used `go.uber.org/fx`)
- [ ] Application Configuration (CLI + Flags (`urfave/cli` or `spf13/cobra`), or `spf13/viper`?)
- [x] Being a web service (`gin`, could have used `net/http` or `chi`)
- [x] Unit Testing (Built-in Go Unit Testing + `stretchr/testify`)
- [x] Logging Config (Built-in `slog`, could have used `go.uber.org/zap`?)
- [ ] Tracing configuration (OpenTelemetry)
- [ ] Metrics configuration (OpenTelemetry + Prometheus Registry/Endpoint)
- [x] Health checks (Custom Liveness and Readiness endpoints)
- [ ] Database migration deployments (`golang-migrate/migrate`)

Build / CI:

- [x] Test & Build automation (Make, Go CLIs `test` & `build`, GitHub Actions)
- [x] Packaging and pushing a container image (Dockerfile, GitHub Actions)
- [ ] Automated publishing of Contract Testing Contracts and Results (PACT Broker, GitHub Actions)
- [x] Lint/Scan/Generate Go code (Make, Go CLIs + `golangci-lint`)
    - `go mod verify` - Ensures dependencies haven't been modified since last downloaded.
    - `go mod tidy` (no diff) - Fixes `go.mod` file to meet requirements for building module (should always be the case)
    - `go fmt` (no diff) - Formats the code according to Go canonical style (should always be the case)
    - `go vet` - Lints the code for common Go mistakes, etc.
    - `golangci-lint` - Run more linters to lint the code.
    - `go generate` (no diff) - Run `//go:generate` scripts to ensure up-to-date generations exist (should always be the
      case)
- [x] Validate MkDocs (`mkdocs build` with `strict` mode)

Deployment / CD:

- Not covered, but would use Kubernetes Deployment (or Argo Rollout, or similar)
- Expected that an external CD system would deploy to Kubernetes (e.g. Argo CD)
- See [init-microservice](https://github.com/harmelodic/init-microservice) for examples.

Infrastructure as Code:

- Not covered, but would use Terraform
- Expected that an external CD system would apply Terraform (e.g. Atlantis)
- See [init-microservice](https://github.com/harmelodic/init-microservice) for examples.

Reference implementation examples (production):

- [x] Application Structure Example (account)
    - Reasonably decoupled layers/components
    - Domain-driven
    - Scoped explicit exception handling
    - Simple reusable model, mapping done in layers (if needed)
    - Dependency Injection used
    - Basic CRUD (as other implementations covered in other reference implementations)
- [x] DB Client (`sqlx`)
    - Could have used `database/sql` but `sqlx` had built-in struct marshalling/row mapping.
    - Could have used `SQLC` but compiling & generating a client from SQL is needless complexity.
    - Could have used `GORM` but I don't like ORMs.
- [ ] HTTP Client (Built-in `net/http`?)

Reference implementations (testing):

- [ ] Provider Contract Testing the Controller (PACT)
- [ ] Consumer Contract Testing the HTTP Client (PACT)
- [x] Integration Testing the Repository (Testcontainers)
- [ ] Integration Testing the Event Publisher (Testcontainers)
- [ ] Integration Testing the Event Subscriber (Testcontainers)

## Documentation

Uses `mkdocs` to handle documentation, which requires Python (hence the `requirements.txt`).

Run docs locally by doing:

```bash
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
mkdocs serve
```

Then open at http://localhost:8000

## Local development

Prerequisites:

- Go is installed
- `golangci-lint` is installed

### Running tests

```shell
make test
```

### Building an executable

```shell
make build
# or just
make
```

### Running the app

```shell
go run ./internal
# or
make build && ./app
```

- Endpoints accessible on http://localhost:8080
- Management endpoints on `/management/...`
