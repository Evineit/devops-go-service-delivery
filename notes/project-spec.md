# Project Specification

This document contains the original requirements for the `devops-go-service-delivery` project.

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

### Requirement 1 - Dockerize the service

Your Dockerfile should:

- use multi-stage build
- compile the Go binary in builder stage
- copy only binary into runtime image
- expose app port
- run as non-root if possible

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

### Requirement 3 - Helm chart

After raw manifests work, convert them into Helm.

Your chart should support:

- image repository/tag
- replica count
- service port
- env variables
- resource limits
- probes

### Requirement 4 - CI pipeline

Use GitHub Actions or GitLab CI.

Pipeline stages should be:

- lint or format check
- test
- build binary
- build image

Optional:

- push image to registry

### Requirement 5 - Troubleshooting practice

You must create at least 5 failure scenarios with symptom, root cause, diagnosis command, and fix.

## Deliverables

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
