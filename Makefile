APP_NAME := bank-backend
GO_MAIN := ./cmd/app
DOCKER_IMAGE := $(APP_NAME):local
K8S_NAMESPACE := bank
K8S_DIR := k8s
KAFKA_TOPIC ?= bank-events
KAFKA_PARTITIONS ?= 1
KAFKA_REPLICATION_FACTOR ?= 1
SWAG_VERSION := v1.16.6
REPOSITORY_DIR := $(abspath ../Bank-repository-service)
EXCHANGE_RATE_DIR := $(abspath ../Bank-exhange-rate-service)

.DEFAULT_GOAL := help

.PHONY: help run build test fmt vet tidy check swagger \
	docker-build docker-up docker-down docker-restart docker-logs docker-ps \
	kafka-up kafka-logs kafka-topics kafka-topic-create \
	k8s-apply k8s-delete k8s-status k8s-logs k8s-port-forward

help: ## Показати список доступних команд
	@awk 'BEGIN {FS = ":.*## "; printf "Usage: make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-22s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Запустити застосунок локально
	go run $(GO_MAIN)

build: ## Зібрати Go-бінарний файл
	go build -trimpath -o $(APP_NAME)$(if $(filter Windows_NT,$(OS)),.exe,) $(GO_MAIN)

test: ## Запустити всі Go-тести
	go test ./...

fmt: ## Відформатувати Go-код
	go fmt ./...

vet: ## Запустити статичну перевірку Go
	go vet ./...

tidy: ## Оновити go.mod і go.sum
	go mod tidy

check: fmt vet test ## Форматування, vet і тести

swagger: ## Generate Swagger documentation
	go run github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION) init -g cmd/app/main.go -o docs --parseDependency --parseInternal

docker-build: ## Зібрати Docker-образ застосунку
	docker build -t $(DOCKER_IMAGE) .

docker-up: ## Запустити застосунок і Kafka через Docker Compose
	docker compose up -d --build

docker-down: ## Зупинити Docker Compose без видалення Kafka-даних
	docker compose down

docker-restart: docker-down docker-up ## Перезапустити Docker Compose

docker-logs: ## Показувати логи всіх Compose-сервісів
	docker compose logs -f

docker-ps: ## Показати стан Compose-сервісів
	docker compose ps

kafka-up: ## Запустити лише Kafka через Docker Compose
	docker compose up -d kafka

kafka-logs: ## Показувати логи Kafka
	docker compose logs -f kafka

kafka-topics: ## Показати список Kafka-топіків
	docker compose exec kafka /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list

kafka-topic-create: ## Створити топік; KAFKA_TOPIC=name make kafka-topic-create
	docker compose exec kafka /opt/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--create --if-not-exists \
		--topic $(KAFKA_TOPIC) \
		--partitions $(KAFKA_PARTITIONS) \
		--replication-factor $(KAFKA_REPLICATION_FACTOR)

k8s-apply: ## Застосувати всі Kubernetes-маніфести
	kubectl apply -f $(K8S_DIR)/namespace.yaml
	kubectl apply -f $(K8S_DIR)/kafka.yaml -f $(K8S_DIR)/app.yaml

k8s-delete: ## Видалити ресурси та namespace з Kubernetes
	kubectl delete -f $(K8S_DIR)/app.yaml -f $(K8S_DIR)/kafka.yaml --ignore-not-found
	kubectl delete -f $(K8S_DIR)/namespace.yaml --ignore-not-found

k8s-status: ## Показати pod, service, deployment і StatefulSet
	kubectl get pods,services,deployments,statefulsets -n $(K8S_NAMESPACE)

k8s-logs: ## Показувати логи bank-backend у Kubernetes
	kubectl logs -f deployment/$(APP_NAME) -n $(K8S_NAMESPACE)

k8s-port-forward: ## Відкрити Kubernetes-застосунок на localhost:8080
	kubectl port-forward service/$(APP_NAME) 8080:8080 -n $(K8S_NAMESPACE)

.PHONY: manual-up
manual-up: ## Start infrastructure and all Go services in separate PowerShell windows
	docker compose down
	docker compose -f "$(REPOSITORY_DIR)/compose.yaml" --profile kafka up -d --wait postgres redis kafka
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(REPOSITORY_DIR)''; go run ./cmd/app'"
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(EXCHANGE_RATE_DIR)''; $$env:KAFKA_BROKERS=''localhost:9092''; go run ./cmd/app'"
	powershell.exe -NoProfile -Command "Start-Process powershell.exe -ArgumentList '-NoExit','-Command','Set-Location ''$(CURDIR)''; $$env:KAFKA_BROKERS=''localhost:9092''; go run ./cmd/app'"

.PHONY: test-cover
test-cover: ## Run tests and keep the coverage profile outside the repository root
	powershell.exe -NoProfile -Command "New-Item -ItemType Directory -Force '.coverage' | Out-Null"
	go test -coverprofile=.coverage/coverage.out ./...
