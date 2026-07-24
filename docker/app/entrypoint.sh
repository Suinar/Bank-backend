#!/bin/bash
set -euo pipefail

KAFKA_CONFIG=/opt/kafka/config/bank-server.properties
KAFKA_DATA=/var/lib/kafka/data
app_pid=""

if [[ ! -f "$KAFKA_DATA/meta.properties" ]]; then
  cluster_id="$(/opt/kafka/bin/kafka-storage.sh random-uuid)"
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

/app/bank-backend &
app_pid=$!

set +e
wait -n "$kafka_pid" "$app_pid"
status=$?
shutdown
exit "$status"
