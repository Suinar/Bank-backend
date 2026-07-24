# Kafka integration

## Implemented in Bank-backend

`Bank-backend` publishes versioned audit events to `bank.backend.audit.v1`
after HTTP requests complete. Events contain operation metadata only:

- event and correlation identifiers;
- producer, event type and schema version;
- HTTP method and route template;
- status, outcome, duration and timestamp.

Request bodies, passwords, card numbers, email addresses and phone numbers are
not published. Kafka delivery is asynchronous, so audit consumers do not add
latency to API responses. The service creates its owned topic idempotently with
three partitions and replication factor three.

Repository-owned domain events are not duplicated here. The repository service
publishes committed `user`, `account`, `card`, `credit`, `deposit`, and
`currency` changes through its transactional outbox.

## Next cross-service phase

The following work requires coordinated changes outside this repository:

1. Add consumers for exchange-rate request topics and define response topics.
2. Define shared request/reply envelopes, timeouts and correlation semantics.
3. Decide which gRPC operations genuinely benefit from asynchronous commands;
   keep queries and operations needing an immediate response on gRPC.
4. Add retry topics, dead-letter topics and idempotent consumer storage.
5. Add contract and end-to-end tests covering all three services.
6. Align exchange-rate topic replication with the three-broker cluster.

Kafka request/reply must not replace gRPC until the corresponding consumers and
failure handling exist; publishing commands without consumers would lose work.
