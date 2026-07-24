# Build a minimal static binary, then copy only runtime essentials into the final image.
FROM golang:1.25.5-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/bank-backend ./cmd/app

FROM apache/kafka:4.0.0

USER root
WORKDIR /app
COPY --from=builder /out/bank-backend ./bank-backend
COPY --chmod=755 docker/app/entrypoint.sh ./entrypoint.sh
COPY docker/kafka/server.properties /opt/kafka/config/bank-server.properties
RUN mkdir -p /var/lib/kafka/data \
    && chown -R appuser:appuser /app /var/lib/kafka/data /opt/kafka/config/bank-server.properties

USER appuser
EXPOSE 8080 9092 29092

ENTRYPOINT ["/app/entrypoint.sh"]
