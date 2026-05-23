.PHONY: run build test fmt docker-build helm-template k8s-apply

APP_NAME := user-service
IMAGE := myrepo/devops-go-service-delivery:latest

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP_NAME) ./cmd/api

test:
	go test ./...

fmt:
	go fmt ./...

docker-build:
	docker build -t $(IMAGE) .

helm-template:
	helm template ./chart --values ./chart/values.yaml

k8s-apply:
	kubectl apply -f k8s/deployment.yaml -f k8s/service.yaml
