# Distributed logging System

A distributed logging system made using gin framework where multiple services can log messages to a central logging server.

### Requirements

- Go version 1.22.4

### Steps

- `git clone https://github.com/asushrestha/distributed-logging-system.git`
- `go mod init`
- `go mod tidy`
- `go run .`

## Viewing the Logs:

Browse http://localhost:8081/logs from your browser.
or

```bash
curl --location 'http://localhost:8081/logs'
```

## Filtering the Logs:

If you want to add filters for the logs, you can add it in the query parameters of the url:
`http://localhost:8081/logs?serviceName={{serviceName}}&severity={{severity}}&startTime=2024-06-23T18:20:39.809144351Z&endTime=2024-06-23T18:25:39.809144351Z`

Severity:

- WARN
- INFO
- ERROR

ServiceName:

- LoggerServiceA
- LoggerServiceB
- LoggerSeriviceC

#Test

- `go test`
