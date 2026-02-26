# api-stress-test
Technical CLI project for HTTP stress testing and execution reporting.



## Challenge
Build a CLI system in Go to run load tests against a web service.
The user must provide the service URL, total number of requests, and the concurrency level.
The system must generate a report with specific metrics after test execution.



## Run manually
``` shell
docker build -t xavierpms/api-stress-test:v1 .
docker run xavierpms/api-stress-test:v1 --url=https://google.com --requests=50 --concurrency=10
```

## Run with make
``` shell
make build
make run
make runc
make push
make verify
```

You can override image and tag when needed:
``` shell
make push IMAGE=rafaelpm/api-stress-test TAG=latest
```

## Run tests
``` shell
go test ./...
```

## Known Targets for Validation
Use deterministic endpoints when validating status-code behavior:

- Always `200`:
``` shell
docker run rafaelpm/api-stress-test:v1 --url=https://httpbin.org/status/200 --requests=100 --concurrency=10
```

- Always `404`:
``` shell
docker run rafaelpm/api-stress-test:v1 --url=https://httpbin.org/status/404 --requests=100 --concurrency=10
```

- Always `500`:
``` shell
docker run rafaelpm/api-stress-test:v1 --url=https://httpbin.org/status/500 --requests=100 --concurrency=10
```

Note: public sites like Google may apply anti-bot protections and not return `200` consistently for automated traffic.



## Language and Runtime Features
- CLI: Cobra
- net/http
- context
- wait groups
- channels
- mutex

## Architecture (Simple and Testable)
- `internal/domain`: domain entities
- `internal/application`: application rules and orchestration (`RunStressTestUseCase`)
- `internal/infrastructure/repository`: infrastructure adapters (HTTP and report generation)
- `internal/infrastructure/cli`: input layer (CLI)

Key files:
- `internal/infrastructure/cli/run_stress_command.go`: CLI command flow and input parsing
- `internal/application/run_stress_workflow.go`: orchestration use case
- `internal/application/execute_requests.go`: HTTP execution use case
- `internal/application/build_report.go`: report generation use case
- `internal/infrastructure/repository/http_request_repository.go`: HTTP adapter
- `internal/infrastructure/repository/status_report_repository.go`: status aggregation adapter
- `internal/domain/repository_contracts.go`: domain contracts for repositories

Applied principles:
- dependency injection to improve unit testability
- error propagation without `log` side effects inside use cases
- clear responsibility boundaries across input parsing, execution, and reporting



## Requirements: Input Parameters
CLI input parameters:
- `--url`: target service URL.
- `--requests`: total number of requests.
- `--concurrency`: number of concurrent workers.

## Requirements: Test Execution
- Send HTTP requests to the provided URL.
- Distribute requests according to the defined concurrency level.
- Ensure the configured total number of requests is executed.

## Requirements: Report
Generate a final report including:
- Total execution time.
- Total number of requests executed.
- Number of requests with HTTP 200 status.
- Distribution of other HTTP status codes (such as 404, 500, etc.).

## Requirements: App Execution
The application must support execution via Docker. Example:
``` shell
docker run <your image> --url=https://google.com --requests=1000 --concurrency=10
```
