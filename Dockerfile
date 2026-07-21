
# Build a minimal static binary, then copy only runtime essentials into the final image.
FROM golang:1.25.5-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/bank-backend ./cmd/app

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app
COPY --from=builder /out/bank-backend ./bank-backend

USER app
EXPOSE 8080

ENTRYPOINT ["/app/bank-backend"]
