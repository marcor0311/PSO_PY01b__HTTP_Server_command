# User Manual – HTTP Server (RedUnix)

## Requirements
- Go
- Git
- Postman (optional, for testing)

## Installation

```bash
git clone https://github.com/marcor0311/PSO_PY01b__HTTP_Server_command.git
cd httpserver
go mod tidy
```

## Running the Server

```bash
cd go run cmd/server/main.go
```

This will start the TCP server on port `8080`:

```
[RedUnix] HTTP Server listening on :8080
```

## How to Use the Server

The server listens for HTTP/1.0 requests over TCP. You can test it with curl or Postman.

Example requests:

```bash
curl "http://localhost:8080/fibonacci?num=10"
curl "http://localhost:8080/random?count=5&min=1&max=10"
curl "http://localhost:8080/reverse?text=hello"
curl "http://localhost:8080/timestamp"
curl "http://localhost:8080/help"
```

## Testing with Postman

1. Open Postman
2. Click "Import" → "File" → Select `server.postman_collection.json`
3. Use the saved examples to test each route

## Available Endpoints

| Route                                   | Description                                |
| --------------------------------------- | ------------------------------------------ |
| `/fibonacci?num=n`                      | Returns the n-th Fibonacci number          |
| `/random?count=n&min=a&max=b`           | Generates random numbers                   |
| `/reverse?text=abc`                     | Reverses the input text                    |
| `/toupper?text=abc`                     | Converts text to uppercase                 |
| `/hash?text=abc`                        | Returns the SHA-256 hash of the input      |
| `/timestamp`                            | Returns current server time (RFC3339)      |
| `/help`                                 | Lists all available endpoints              |
| `/status`                               | Server status, uptime, request log         |
| `/createfile?name=x&content=y&repeat=z` | Creates a file                             |
| `/deletefile?name=x`                    | Deletes a file                             |
| `/simulate?seconds=s&task=x`            | Simulates blocking task                    |
| `/sleep?seconds=n`                      | Sleeps the server                          |
| `/loadtest?tasks=n&sleep=s`             | Simulate concurrent requests               |


## Test Coverage Report

```bash
go test -cover ./...
```

# App Architecture 
```
httpserver/
│   └── go.mod
│   └── MANUAL.md
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── constants/
│   │   └── constants.go
│   ├── handlers/
│   │   └── files.go
│   │   └── files_test.go
│   │   └── help.go
│   │   └── help_test.go
│   │   └── math.go
│   │   └── math_test.go
│   │   └── simulate.go
│   │   └── simulate_test.go
│   │   └── status.go
│   │   └── status_test.go
│   │   └── strings.go
│   │   └── strings_test.go
│   │   └── system.go
│   │   └── system_test.go
│   ├── router/
│   │   └── handlers.go
│   │   └── router.go
│   ├── tcp/
│   │   └── client.go
│   │   └── connection.go
│   ├── utils/
        └── http_test.go
│   │   └── http.go
└──  postman/
    └── server.postman_collection.json

```