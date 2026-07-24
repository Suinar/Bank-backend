# Docker configurations

All Docker-related files live in this directory. The complete application stack
is defined in `docker-compose.yml`, while component-specific files have their
own subdirectory.

## Application

The application image files are located in `app/`:

- `app/Dockerfile`
- `app/Dockerfile.dockerignore`

`Dockerfile.dockerignore` uses Docker's Dockerfile-specific ignore-file naming,
so it still applies when the build context is the repository root.

## Kafka

Kafka is not a separate container. Its KRaft properties are defined by
`KAFKA_SERVER_PROPERTIES` in `docker-compose.yml`, and Kafka runs alongside the
Go application inside the `bank-backend-service` container. Start or rebuild
that container with:

```sh
make kafka-up
```

Host applications use `localhost:39092,localhost:39093,localhost:39094`.
Services in the complete Docker Compose stack use
`app:9092,kafka-2:9092,kafka-3:9092`.

The application is exposed on `localhost:18080`. Override the service-specific
host ports with `BANK_BACKEND_HTTP_PORT` and `BANK_BACKEND_KAFKA_PORT`.
Set `BANK_BACKEND_START_APPLICATION=false` when Kafka must run without the Go
application and its gRPC dependencies.

Because Kafka and the application share one container, this stops both:

```sh
make kafka-down
```
