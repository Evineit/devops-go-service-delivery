# devops-go-service-delivery

A small Go HTTP API taken through the full deployment lifecycle: Kubernetes, CI/CD, Helm

## Goal

Build a small Go HTTP API and take it through the full deployment lifecycle:

- write code in Go
- test it
- containerize it
- deploy it to Kubernetes
- package it with Helm
- automate it with CI
- debug failures

## Functional requirements

Your service should expose these endpoints:

### 1. Health endpoint

`GET /health`

Response:

```json
{
  "status": "ok"
}
```

Purpose:

- liveness/readiness
- simple HTTP handler
- health checks in Kubernetes

### 2. List users

`GET /users`

Response:

```json
[
  {
    "id": 1,
    "name": "Kevin",
    "email": "kevin@example.com"
  }
]
```

Purpose:

- JSON response
- slice handling
- route handler basics

### 3. Create user

`POST /users`

Request body:

```json
{
  "name": "Ana",
  "email": "ana@example.com"
}
```

Response:

```json
{
  "id": 2,
  "name": "Ana",
  "email": "ana@example.com"
}
```

Purpose:

- decode JSON body
- validate input
- return created object

### 4. Metrics endpoint

`GET /metrics`

For the first version, this can be simple:

- request counter in memory
- or Prometheus metrics later as stretch goal

Purpose:

- observability mindset
- useful for future projects

## Non-functional requirements

These matter a lot because they make it feel closer to a real service.

The app must:

- run locally with `go run`
- build with `go build`
- run inside Docker
- expose configurable port via environment variable
- log every incoming request
- return proper HTTP status codes
- handle invalid JSON gracefully
- include at least basic tests
- include a Dockerfile
- include Kubernetes manifests
- include a Helm chart
- include a CI pipeline

## Explicit DevOps requirements

Now the DevOps side.

### Requirement 1 - Dockerize the service

Your Dockerfile should:

- use multi-stage build
- compile the Go binary in builder stage
- copy only binary into runtime image
- expose app port
- run as non-root if possible

What you learn:

- smaller images
- build/runtime separation
- production-ish container practices

### Requirement 2 - Kubernetes manifests

You should first create raw manifests:

- `deployment.yaml`
- `service.yaml`

Deployment should include:

- image
- replicas
- env vars
- container port
- liveness probe
- readiness probe
- resource requests/limits

Service should expose:

- app internally in cluster

What you learn:

- direct Kubernetes basics before abstracting with Helm

### Requirement 3 - Helm chart

After raw manifests work, convert them into Helm.

Your chart should support:

- image repository/tag
- replica count
- service port
- env variables
- resource limits
- probes

What you learn:

- templating
- values management
- reusable deployments

### Requirement 4 - CI pipeline

Use GitHub Actions or GitLab CI.

Pipeline stages should be:

- lint or format check
- test
- build binary
- build image

Optional:

- push image to registry

What you learn:

- automation of validation and packaging
- how app code connects to delivery systems

### Requirement 5 - Troubleshooting practice

You should deliberately break the system and fix it.

You must create at least 5 failure scenarios.

Example failures:

- wrong container port
- bad readiness path
- invalid Helm value
- app crashes on startup
- failed JSON request handling
- image tag mismatch

For each one, write:

- symptom
- root cause
- command used to diagnose
- fix

What you learn:

- confidence under failure
- interview stories
- platform engineering mindset

## Deliverables

By the end of Project 0, you should have these artifacts.

### Code deliverables

- working Go API
- clean repo structure
- tests
- Dockerfile
- Makefile
- Kubernetes manifests
- Helm chart
- CI pipeline config

### Documentation deliverables

- README with setup steps
- architecture overview
- endpoint documentation
- troubleshooting notes for at least 5 incidents


## Implementation checklist

This is the "done means done" version.

### Application

- [x] GET /health works
- [x] GET /users works
- [x] POST /users works
- [x] invalid JSON returns 400
- [x] missing fields return 400
- [x] requests are logged
- [x] port is configurable via env var

### Go learning

- [x] used structs
- [x] used package separation
- [x] used JSON encode/decode
- [x] used env config
- [x] wrote unit tests
- [x] built compiled binary

### Container

- [x] Dockerfile works
- [x] app runs in container
- [x] multi-stage build used

### Kubernetes

- [x] deployment works on kind/k3d
- [x] service exposes app
- [x] liveness probe works
- [x] readiness probe works

### Helm

- [x] app deploys with Helm
- [x] values can override image tag and replica count

### CI

- [x] run tests automatically
- [x] run image build automatically

### Troubleshooting

- [x] documented at least 5 break/fix cases

## Quick Start

Run locally:

```bash
go run ./cmd/api
# server on http://localhost:8080
```

Build binary:

```bash
make build
```

Run with Docker (multi-stage build included):

```bash
make docker-build
docker run -p 8080:8080 myrepo/devops-go-service-delivery:local
```

Endpoints:

- `GET /health` -> {"status":"ok"}
- `GET /users` -> list users
- `POST /users` -> create user (JSON body: name, email)
- `GET /metrics` -> {"requests": <count>}

Kubernetes (raw manifests):

```bash
kubectl apply -f k8s/deployment.yaml -f k8s/service.yaml
```

Helm (render templates):

```bash
helm template ./chart --values ./chart/values.yaml
```

CI:

This repo includes a GitHub Actions workflow at `.github/workflows/ci.yml` that runs `go fmt`, `go test`, builds the binary, and builds a Docker image.
