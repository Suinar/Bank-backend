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

Kafka is not a separate container. Its KRaft configuration is stored in
`kafka/server.properties`, and Kafka runs alongside the Go application inside
the `bank-backend` container. Start or rebuild that container with:

```sh
make kafka-up
```

Host applications connect to `localhost:29092`. Services in the complete
Docker Compose stack connect to `app:9092`.

Because Kafka and the application share one container, this stops both:

```sh
make kafka-down
```
