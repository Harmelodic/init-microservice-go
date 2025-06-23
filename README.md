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

Application configuration:

- [ ] Dependency Injection (Plain old root DI, or `go.uber.org/fx`?)
- [ ] Application Configuration (CLI + Flags (`urfave/cli` or `spf13/cobra`), or `spf13/viper`?)
- [ ] Being a web service (`gin` or `net/http` or `chi`?)
- [ ] Unit Testing (Built-in Go Unit Testing + `stretchr/testify`)
- [ ] Logging Config (Built-in `slog` or `go.uber.org/zap`?)
- [ ] Tracing configuration (OpenTelemetry)
- [ ] Metrics configuration (OpenTelemetry + Prometheus Registry/Endpoint)
- [ ] Health checks (Custom?)

Build / CI:

- [x] Test & Build automation (Go CLIs `test` & `build`, GitHub Actions)
- [ ] Packaging and pushing a container image (Dockerfile, GitHub Actions)
- [ ] Automated publishing of Contract Testing Contracts and Results (PACT Broker, GitHub Actions)
- [x] Lint/Scan/Generate Go code (Go CLIs + `golangci-lint`)
  - `go mod verify` - Ensures dependencies haven't been modified since last downloaded.
  - `go mod tidy` (no diff) - Fixes `go.mod` file to meet requirements for building module (should always be the case)
  - `go fmt` (no diff) - Formats the code according to Go canonical style (should always be the case)
  - `go vet` - Lints the code for common Go mistakes, etc.
  - `golangci-lint` - Run more linters to lint the code.
  - `go generate` (no diff) - Run `//go:generate` scripts to ensure up-to-date generations exist (should always be the
    case)
- [x] Validate MkDocs (`mkdocs build` with `strict` mode)

Deployment / CD:

- Kubernetes Deployment (or Argo Rollout, or similar)
- Expected that an external CD system would deploy to Kubernetes (e.g. Argo CD)
- See [init-microservice](https://github.com/harmelodic/init-microservice) for examples.

Infrastructure as Code:

- Terraform
- Expected that an external CD system would apply Terraform (e.g. Atlantis)
- See [init-microservice](https://github.com/harmelodic/init-microservice) for examples.

Reference implementation examples (production):

- [ ] Application Structure Example (account)
  - Reasonably decoupled layers/components
  - Domain-driven
  - Scoped explicit exception handling
  - Simple reusable model, mapping done in layers (if needed)
  - Dependency Injection used
  - No implementation details (as implementations covered in other reference implementations)
- [ ] DB Client (GORM? SQLC? SQLX?)
- [ ] HTTP Client (Built-in HTTP)

Reference implementations (testing):

- [ ] Provider Contract Testing the Controller (PACT)
- [ ] Consumer Contract Testing the HTTP Client (PACT)
- [ ] Integration Testing the Repository (in-memory DB)
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

## Running the app locally

```bash
# TODO
```

- Endpoints accessible on ??? TODO
