APP_NAME := bank-backend
GO_MAIN := ./cmd/app
DOCKER_IMAGE := $(APP_NAME):local
DOCKER_COMPOSE := docker compose -f docker/docker-compose.yml
K8S_NAMESPACE := bank
K8S_DIR := docker/kubernetes
KAFKA_TOPIC ?=
KAFKA_PARTITIONS ?= 3
KAFKA_REPLICATION_FACTOR ?= 1
SWAG_VERSION := v1.16.6
REPOSITORY_DIR := $(abspath ../Bank-repository-service)
EXCHANGE_RATE_DIR := $(abspath ../Bank-exhange-rate-service)

.DEFAULT_GOAL := help

.PHONY: help run build test fmt vet tidy check swagger \
	docker-build docker-up docker-down docker-restart docker-logs docker-ps \
	kafka-up kafka-down kafka-logs kafka-topics kafka-topic-create \
	k8s-apply k8s-delete k8s-status k8s-logs k8s-port-forward

help: ## Show available commands
	@awk 'BEGIN {FS = ":.*## "; printf "Usage: make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-22s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application locally
	go run $(GO_MAIN)

build: ## Build the Go binary
	go build -trimpath -o $(APP_NAME)$(if $(filter Windows_NT,$(OS)),.exe,) $(GO_MAIN)

test: ## Run all Go tests
	go test ./...

fmt: ## Format Go code
	go fmt ./...

vet: ## Run Go static analysis
	go vet ./...

tidy: ## Update go.mod and go.sum
	go mod tidy

check: fmt vet test ## Format, vet, and test

swagger: ## Generate Swagger documentation
	go run github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION) init -g cmd/app/main.go -o docs --parseDependency --parseInternal

docker-build: ## Build the application Docker image
	docker build -f docker/app/Dockerfile -t $(DOCKER_IMAGE) .

docker-up: ## Start the application and Kafka with Docker Compose
	$(DOCKER_COMPOSE) up -d --build

docker-down: ## Stop Docker Compose without deleting Kafka data
	$(DOCKER_COMPOSE) down

docker-restart: docker-down docker-up ## Restart Docker Compose

docker-logs: ## Follow logs from all Compose services
	$(DOCKER_COMPOSE) logs -f

docker-ps: ## Show Compose service status
	$(DOCKER_COMPOSE) ps

kafka-up: ## Start only Kafka with Docker Compose
	BANK_BACKEND_START_APPLICATION=false $(DOCKER_COMPOSE) up -d --build --no-deps app

kafka-down: ## Stop the Kafka container without deleting data
	$(DOCKER_COMPOSE) stop app

kafka-logs: ## Follow Kafka logs
	$(DOCKER_COMPOSE) logs -f app

kafka-topics: ## List Kafka topics
	$(DOCKER_COMPOSE) exec app /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list

kafka-topic-create: ## Create a topic; KAFKA_TOPIC=name make kafka-topic-create
	@test -n "$(KAFKA_TOPIC)" || (echo "KAFKA_TOPIC is required"; exit 1)
	$(DOCKER_COMPOSE) exec app /opt/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--create --if-not-exists \
		--topic $(KAFKA_TOPIC) \
		--partitions $(KAFKA_PARTITIONS) \
		--replication-factor $(KAFKA_REPLICATION_FACTOR)

k8s-apply: ## Apply all Kubernetes manifests
	kubectl apply -f $(K8S_DIR)

k8s-delete: ## Delete Kubernetes resources and namespace
	kubectl delete -f $(K8S_DIR) --ignore-not-found

k8s-status: ## Show pods, services, deployments, and StatefulSets
	kubectl get pods,services,deployments,statefulsets -n $(K8S_NAMESPACE)

k8s-logs: ## Follow bank-backend logs in Kubernetes
	kubectl logs -f deployment/$(APP_NAME) -n $(K8S_NAMESPACE)

k8s-port-forward: ## Expose the Kubernetes application on localhost:8080
	kubectl port-forward service/$(APP_NAME) 8080:8080 -n $(K8S_NAMESPACE)

.PHONY: manual-up
manual-up: ## Start infrastructure and all Go services in separate PowerShell windows
	$(DOCKER_COMPOSE) down
	docker compose -f "$(REPOSITORY_DIR)/compose.yaml" --profile kafka up -d --wait postgres redis kafka
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(REPOSITORY_DIR)''; go run ./cmd/app'"
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(EXCHANGE_RATE_DIR)''; $$env:KAFKA_BROKERS=''localhost:9092''; go run ./cmd/app'"
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(CURDIR)''; $$env:KAFKA_BROKERS=''localhost:9092''; go run ./cmd/app'"

.PHONY: test-cover
test-cover: ## Run tests and keep the coverage profile outside the repository root
	powershell.exe -NoProfile -Command "New-Item -ItemType Directory -Force '.coverage' | Out-Null"
	go test -coverprofile=.coverage/coverage.out ./...

.PHONY: postman-test
postman-test: ## Run Postman integration tests against the local Docker stack
	powershell.exe -NoProfile -ExecutionPolicy Bypass -File test/postman/run.ps1
