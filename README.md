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

- [ ] GET /health works
- [ ] GET /users works
- [ ] POST /users works
- [ ] invalid JSON returns 400
- [ ] missing fields return 400
- [ ] requests are logged
- [ ] port is configurable via env var

### Go learning

- [ ] used structs
- [ ] used package separation
- [ ] used JSON encode/decode
- [ ] used env config
- [ ] wrote unit tests
- [ ] built compiled binary

### Container

- [ ] Dockerfile works
- [ ] app runs in container
- [ ] multi-stage build used

### Kubernetes

- [ ] deployment works on kind/k3d
- [ ] service exposes app
- [ ] liveness probe works
- [ ] readiness probe works

### Helm

- [ ] app deploys with Helm
- [ ] values can override image tag and replica count

### CI

- [ ] run tests automatically
- [ ] run image build automatically

### Troubleshooting

- [ ] documented at least 5 break/fix cases