# Postman integration tests

The collection runs a repeatable Bank-backend workflow:

1. checks service health;
2. verifies invalid input handling;
3. creates a user, currency, account, card, credit and deposit;
4. verifies reads and mutations;
5. removes all created resources;
6. confirms that a removed resource returns `404`.

Every API request checks its status, JSON response, response time and
`X-Correlation-ID` propagation. Test data is generated per run, and no secrets
are stored in the environment file.

## Prerequisites

- the Docker Compose stack is healthy;
- Node.js with `npx` is installed.

Start the stack and run the tests:

```powershell
docker compose -f docker/docker-compose.yml up -d --build --wait
.\test\postman\run.ps1
```

Override the target URL when necessary:

```powershell
.\test\postman\run.ps1 -BaseUrl http://localhost:8080
```

You can also import both JSON files into Postman and run the collection with
the `Bank Backend - Local` environment.

The pinned Newman version makes local and CI results reproducible. Cleanup
requests still run after an earlier assertion failure, and Newman exits with a
non-zero status when the run contains failures.
