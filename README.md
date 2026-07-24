# Bank Backend

HTTP API gateway and application service for the banking platform. It exposes
user, account, card, credit, deposit, currency, and exchange-rate operations,
delegates persistence and exchange-rate queries to gRPC services, and publishes
non-sensitive HTTP audit events to Kafka.

## How the Service Works

1. The application loads configuration from environment variables and an
   optional local `.env` file.
2. It connects to the repository and exchange-rate services over gRPC.
3. When Kafka is enabled, it ensures that the backend-owned audit topic exists
   and initializes an asynchronous producer.
4. Gin exposes the public HTTP API and Swagger UI.
5. Application services validate requests and call the appropriate gRPC
   repositories.
6. Completed HTTP operations are published to Kafka with their correlation ID,
   route, status, outcome, duration, and timestamp.
7. On `SIGINT` or `SIGTERM`, the HTTP server, Kafka writer, and gRPC connections
   shut down gracefully.

Business requests that require an immediate response use HTTP and gRPC. Kafka is
used for asynchronous audit and domain events, not as a replacement for every
synchronous method call.

## Main Components

- Go 1.25.5
- Gin HTTP framework
- gRPC and Protocol Buffers
- Apache Kafka 4.0 in KRaft mode
- PostgreSQL 16
- Redis 7
- Swagger/OpenAPI
- Docker and Docker Compose
- Kubernetes
- Viper
- GoMock
- Postman and Newman

PostgreSQL and Redis are owned by the repository and exchange-rate services but
are included in the root Docker Compose stack for complete local development.

## Features

- REST API for seven banking resource groups
- gRPC clients for repository and exchange-rate services
- Environment-first configuration with optional `.env`
- Three-broker Kafka KRaft cluster
- Versioned Kafka audit events
- Request correlation through `X-Correlation-ID`
- Automatic Kafka topic initialization
- Swagger UI and generated OpenAPI documentation
- Health endpoint for containers and Kubernetes
- Graceful shutdown
- Unit tests and Postman integration tests
- Docker Compose full-stack environment
- Kubernetes StatefulSet, Deployment, Services, HPA, and disruption budgets

## Project Structure

```text
.
|-- cmd/app/                         # Application entry point
|-- docs/                            # Swagger output and Kafka documentation
|-- internal/
|   |-- brokers/kafka/               # Kafka topic and audit publisher
|   |-- configs/                     # Environment and .env configuration
|   |-- delivery/
|   |   |-- grps/client/             # Outbound gRPC clients
|   |   `-- http/
|   |       |-- handler/             # Gin handlers and routes
|   |       `-- middleware/          # Kafka audit middleware
|   |-- logging/                     # Layer-specific logging
|   |-- mocks/                       # Generated GoMock implementations
|   |-- service/                     # Application services and mapping
|   `-- test/                        # Shared constants and fixtures
|-- docker/
|   |-- app/                         # Multi-stage application image
|   |-- kubernetes/                  # Kubernetes manifests
|   `-- docker-compose.yml           # Complete local stack
|-- test/postman/                    # Postman collection and Newman runner
|-- .env.example                     # Configuration template
|-- Makefile                         # Development commands
`-- README.md
```

## Requirements

For the complete Docker environment:

- Docker Desktop or Docker Engine;
- Docker Compose;
- local sibling checkouts of:
  - `../Bank-repository-service`;
  - `../Bank-exhange-rate-service`.

For local development:

- Go 1.25.5 or a compatible newer version;
- running repository and exchange-rate gRPC services;
- a reachable Kafka cluster when Kafka publishing is enabled;
- GNU Make or compatible tooling, optionally;
- Node.js with `npx` for Newman integration tests.

For Kubernetes:

- a Kubernetes cluster;
- `kubectl`;
- locally available or registry-hosted images for all three bank services.

## Configuration

Create a local configuration file:

```bash
cp .env.example .env
```

On Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

The `.env` file is optional. Process environment variables override file
values, which allows the same image to run locally, in Docker, Kubernetes, and
CI.

### Application and dependencies

| Variable | Default value | Purpose |
|---|---:|---|
| `APP_ENV` | `local` | Runtime environment name |
| `HTTP_HOST` | `0.0.0.0` | HTTP bind address |
| `HTTP_PORT` | `8080` | HTTP server port |
| `BANK_BACKEND_HTTP_PORT` | `18080` | Docker host HTTP port |
| `GRPC_REPOSITORY_URL` | `localhost:50052` | Repository service address |
| `GRPC_EXCHANGE_RATE_URL` | `localhost:50053` | Exchange-rate service address |
| `CARD_BIN` | `537541` | Six-digit card issuer identifier |
| `BANK_BACKEND_START_APPLICATION` | `true` | Starts the backend process in the combined Docker image |

### Kafka

| Variable | Default value | Purpose |
|---|---:|---|
| `KAFKA_ENABLED` | `true` | Enables backend audit publishing |
| `KAFKA_BROKERS` | `localhost:39092,localhost:39093,localhost:39094` | Comma-separated bootstrap brokers |
| `KAFKA_CLIENT_ID` | `bank-backend` | Producer name stored in event envelopes |
| `KAFKA_AUDIT_TOPIC` | `bank.backend.audit.v1` | Backend-owned audit topic |
| `KAFKA_TOPIC_PARTITIONS` | `3` | Audit topic partition count |
| `KAFKA_TOPIC_REPLICATION_FACTOR` | `3` | Audit topic replication factor |
| `BANK_BACKEND_KAFKA_PORT` | `39092` | First Docker broker host port |
| `BANK_BACKEND_KAFKA_2_PORT` | `39093` | Second Docker broker host port |
| `BANK_BACKEND_KAFKA_3_PORT` | `39094` | Third Docker broker host port |

The replication factor must not exceed the number of available brokers. Keep
credentials and future Kafka TLS/SASL values in secrets, never in Git.

## Running with Docker

Build all images and start the complete stack:

```bash
docker compose -f docker/docker-compose.yml up -d --build --wait
```

Alternatively:

```bash
make docker-up
```

The stack includes:

- Bank backend;
- repository service;
- exchange-rate service;
- three Kafka brokers;
- PostgreSQL;
- Redis.

Check status and logs:

```bash
make docker-ps
make docker-logs
```

Stop the environment without deleting persistent volumes:

```bash
make docker-down
```

Default host endpoints:

| Component | Address |
|---|---|
| Bank Backend API | `http://localhost:18080` |
| Swagger UI | `http://localhost:18080/swagger/index.html` |
| Kafka broker 1 | `localhost:39092` |
| Kafka broker 2 | `localhost:39093` |
| Kafka broker 3 | `localhost:39094` |

## Running Locally

Start the external dependencies first and copy `.env.example` to `.env`.
Then run:

```bash
go run ./cmd/app
```

Or:

```bash
make run
```

For a local backend process connecting to the Docker infrastructure, use the
host gRPC and Kafka ports in `.env`.

## Make Commands

Run `make help` to display the complete list.

| Command | Description |
|---|---|
| `make run` | Run the service locally |
| `make build` | Build the Go binary |
| `make test` | Run all Go tests |
| `make test-cover` | Generate the coverage profile |
| `make fmt` | Format Go code |
| `make vet` | Run static analysis |
| `make check` | Format, vet, and test |
| `make swagger` | Regenerate Swagger documentation |
| `make docker-build` | Build `bank-backend:local` |
| `make docker-up` | Build and start the complete Compose stack |
| `make docker-down` | Stop the Compose stack |
| `make docker-logs` | Follow all Compose logs |
| `make docker-ps` | Show Compose service status |
| `make kafka-topics` | List Kafka topics |
| `make kafka-topic-create` | Create a configured Kafka topic |
| `make k8s-apply` | Apply Kubernetes manifests |
| `make k8s-delete` | Delete Kubernetes resources |
| `make postman-test` | Run Postman integration tests |

## HTTP API

The base path is:

```text
/api/v1
```

### Resource groups

| Resource | Main operations |
|---|---|
| Users | List, get by ID/email/phone, create, update, change password, delete |
| Accounts | List, get by ID/user, create, update, block, close, delete |
| Cards | List, get by ID/user/number, create, block, delete |
| Credits | List, get by ID/user, create, repay, delete |
| Deposits | List, get by ID/user, create, replenish, delete |
| Currencies | List, get by ID/ISO/symbol, create, update, delete |
| Exchange rates | Get one relative rate or all rates for a source currency |

### Utility endpoints

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/health` | Container and Kubernetes health check |
| `GET` | `/swagger/index.html` | Interactive API documentation |

All API operations return JSON. Successful operations currently return
`200 OK`; validation, missing resources, and unexpected failures map to
`400`, `404`, and `500`.

## Swagger

Open Swagger UI after starting the service:

```text
http://localhost:18080/swagger/index.html
```

Regenerate the OpenAPI files after changing handler annotations:

```bash
make swagger
```

Generated files are stored in `docs/`.

## Kafka

The backend creates and publishes to:

```text
bank.backend.audit.v1
```

The event envelope includes:

```json
{
  "event_id": "c03f3a9d2a27418ca125ad005c1c530b",
  "event_type": "backend.http.operation.completed",
  "event_version": 1,
  "producer": "bank-backend",
  "correlation_id": "request-123",
  "method": "GET",
  "route": "/api/v1/currencies",
  "status_code": 200,
  "outcome": "success",
  "duration_ms": 2,
  "occurred_at": "2026-07-24T20:28:04Z"
}
```

Request bodies, passwords, card numbers, email addresses, and phone numbers are
not published. Delivery is asynchronous and does not add Kafka consumer latency
to HTTP responses.

Domain events for users, accounts, cards, credits, deposits, and currencies are
owned by the repository service and published through its transactional outbox.
The backend does not duplicate those events.

See `docs/kafka-integration.md` for the current cross-service integration
boundary and future request/reply work.

## Tests

Run the Go test suite:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Run all standard checks:

```bash
make check
```

Generate coverage:

```bash
make test-cover
go tool cover -func=.coverage/coverage.out
```

### Postman integration tests

The Postman suite performs a complete isolated workflow:

- health and validation checks;
- creation of users, currencies, accounts, cards, credits, and deposits;
- reads, updates, repayment, replenishment, and card blocking;
- cleanup of all created resources;
- status, JSON, response-time, and correlation-ID assertions.

Run it against the local Docker stack:

```powershell
.\test\postman\run.ps1
```

Or:

```bash
make postman-test
```

The collection, environment, runner, and detailed instructions are under
`test/postman/`.

## Architecture

```text
HTTP clients
     |
     v
Gin handlers and middleware
     |
     v
Application services
     |
     +---------------------> Repository service over gRPC
     |
     +---------------------> Exchange-rate service over gRPC
     |
     `---------------------> Kafka audit producer

Repository service
     |
     +---------------------> PostgreSQL
     +---------------------> Redis
     `---------------------> Kafka transactional outbox

Exchange-rate service
     |
     +---------------------> Monobank API
     +---------------------> Redis
     `---------------------> Kafka topic infrastructure
```

## Kubernetes

Kubernetes manifests are stored in `docker/kubernetes` and provide:

- the shared `bank` namespace;
- a three-node Kafka KRaft StatefulSet;
- persistent Kafka volumes;
- headless and client Kafka Services;
- a two-replica backend Deployment;
- startup, readiness, and liveness probes;
- resource requests and limits;
- non-root security contexts;
- a HorizontalPodAutoscaler;
- PodDisruptionBudgets for the backend and Kafka.

Apply the shared backend infrastructure first, followed by the dependent
services:

```bash
kubectl apply -f docker/kubernetes
kubectl apply -f ../Bank-repository-service/docker/kubernetes
kubectl apply -f ../Bank-exhange-rate-service/docker/kubernetes
```

Check the rollout:

```bash
kubectl get pods,services -n bank
kubectl rollout status statefulset/kafka -n bank
kubectl rollout status deployment/bank-backend -n bank
```

Expose the API locally:

```bash
kubectl port-forward -n bank service/bank-backend 18080:8080
```

Replace local image tags with immutable registry tags and configure production
StorageClasses, Secrets, NetworkPolicies, and Kafka TLS/SASL before production
deployment.

## Docker Image

Build the image manually:

```bash
docker build -f docker/app/Dockerfile -t bank-backend:local .
```

The multi-stage image contains the static Go binary and the Kafka runtime used
by the local combined broker/backend container. In Kubernetes,
`START_KAFKA=false` runs only the backend process because Kafka is deployed as a
separate StatefulSet.
