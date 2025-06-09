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
- [ ] Auto-updates (Renovate)
- [ ] Dependency Management (go.mod)

Application configuration:

- [ ] Dependency Injection and Application Configuration
- [ ] Unit Testing (Built-in Go Unit Testing)
- [ ] Logging Config (???)
- [ ] Tracing configuration (??? + OpenTelemetry)
- [ ] Metrics configuration (??? + Prometheus Registry/Endpoint)

Build / CI:

- [ ] Test & Build automation (Go `build` CLI, GitHub Actions)
- [ ] Packaging and pushing a container image (???, GitHub Actions)
- [ ] Automated publishing of Contract Testing Contracts and Results (PACT Broker, GitHub Actions)
- [ ] Lint/Scan Java code (Go `vet` CLI)
- [ ] Validate MkDocs (`mkdocs build` with `strict` mode)

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
