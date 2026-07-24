#!/bin/bash
set -euo pipefail

if [[ "${START_KAFKA:-true}" != "true" ]]; then
  exec /app/bank-backend
fi

KAFKA_CONFIG=/tmp/bank-backend-kafka.properties
KAFKA_DATA=/var/lib/kafka/data
app_pid=""

: "${KAFKA_SERVER_PROPERTIES:?KAFKA_SERVER_PROPERTIES must be set}"
printf '%s\n' "$KAFKA_SERVER_PROPERTIES" > "$KAFKA_CONFIG"

if [[ ! -f "$KAFKA_DATA/meta.properties" ]]; then
  cluster_id="${KAFKA_CLUSTER_ID:-$(/opt/kafka/bin/kafka-storage.sh random-uuid)}"
  /opt/kafka/bin/kafka-storage.sh format --cluster-id "$cluster_id" --config "$KAFKA_CONFIG"
fi

/opt/kafka/bin/kafka-server-start.sh "$KAFKA_CONFIG" &
kafka_pid=$!

shutdown() {
  kill -TERM "$kafka_pid" 2>/dev/null || true
  if [[ -n "$app_pid" ]]; then
    kill -TERM "$app_pid" 2>/dev/null || true
    wait "$app_pid" 2>/dev/null || true
  fi
  wait "$kafka_pid" 2>/dev/null || true
}
trap shutdown SIGINT SIGTERM

for _ in {1..60}; do
  if /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list >/dev/null 2>&1; then
    break
  fi
  if ! kill -0 "$kafka_pid" 2>/dev/null; then
    wait "$kafka_pid"
    exit 1
  fi
  sleep 1
done

if ! /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list >/dev/null 2>&1; then
  echo "Kafka did not become ready within 60 seconds" >&2
  exit 1
fi

if [[ "${START_APPLICATION:-true}" != "true" ]]; then
  wait "$kafka_pid"
  exit $?
fi

set +e
while kill -0 "$kafka_pid" 2>/dev/null; do
  /app/bank-backend &
  app_pid=$!
  wait "$app_pid"
  app_status=$?
  app_pid=""

  if ! kill -0 "$kafka_pid" 2>/dev/null; then
    break
  fi

  echo "Application exited with status $app_status; restarting in 2 seconds" >&2
  sleep 2
done

wait "$kafka_pid"
status=$?
shutdown
exit "$status"
