# devops-go-service-delivery

A Go HTTP API taken through the full deployment lifecycle — containerized with Docker, deployed to Kubernetes, packaged with Helm, and automated with GitHub Actions CI.

## Architecture

```
cmd/api/main.go              Entry point, route registration, server start
internal/handlers/handlers.go HTTP handlers for all endpoints
internal/store/database.go    In-memory data store with mutex-guarded concurrent access
internal/middleware/logging.go Request logging middleware
internal/metrics/metrics.go   Atomic request counter
```

The app uses Go's standard `http.ServeMux` with a logging middleware wrapper. All data is stored in-memory with `sync.RWMutex` for concurrency safety. The metrics package provides an atomic request counter exposed at `/metrics`.

## API Endpoints

| Method | Path                | Description                          | Response                         |
|--------|---------------------|--------------------------------------|----------------------------------|
| GET    | `/health`           | Liveness/readiness check             | `{"status": "ok"}`              |
| GET    | `/users`            | List all users                       | Array of user objects            |
| POST   | `/users`            | Create a user                        | Created user object (201)        |
| GET    | `/user/profile?clientId={id}` | Get a single user by ID | User object                       |
| PATCH  | `/user/profile?clientId={id}` | Update a user by ID    | Updated user object               |
| GET    | `/metrics`          | Request counter                      | `{"requests": <count>}`         |
| GET    | `/`                 | Welcome message                      | Plain text                       |

## Quick Start

```bash
# Run locally
go run ./cmd/api

# Build binary
make build

# Run tests
make test

# Run with Docker
make docker-build
docker run -p 8080:8080 myrepo/devops-go-service-delivery:latest

# Deploy to Kubernetes
kubectl apply -f k8s/deployment.yaml -f k8s/service.yaml

# Render Helm templates
helm template ./chart --values ./chart/values.yaml
```

The server port defaults to `8080` and is configurable via the `PORT` environment variable.

## DevOps Stack

| Component    | Implementation                                             |
|-------------|------------------------------------------------------------|
| Container   | Multi-stage Docker build (golang:1.26.1-alpine → alpine:3.18), non-root user |
| Kubernetes  | Raw manifests + Helm chart with configurable values        |
| CI/CD       | GitHub Actions: format check, tests, binary build, Docker image build |
| Deployment  | Deployed to local k3d cluster for testing                  |
| Monitoring  | Request logging via `slog`, in-memory metrics counter      |

## Troubleshooting

Five documented failure scenarios are available in [notes/troubleshooting.md](notes/troubleshooting.md), covering wrong container port, bad readiness path, invalid Helm values, app crash on startup, and invalid JSON handling. Each includes symptom, root cause, diagnosis command, and fix.

## Design Spec

The original project requirements and implementation checklist are archived in [notes/project-spec.md](notes/project-spec.md).
