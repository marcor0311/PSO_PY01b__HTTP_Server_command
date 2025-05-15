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

## Running Unit Tests

```bash
go test -v ./...
```

## Test Coverage Report

```bash
go test -cover ./...
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

# App Architecture 
```
httpserver/
├── main.go                  // Entry point, starts TCP server
├── internal/
│   ├── constants/
│   │   └── constants.go     // Route constants and common values
│   ├── handlers/
│   │   ├── fibonacci.go     // Handler for /fibonacci
│   │   ├── random.go        // Handler for /random
│   │   ├── reverse.go       // Handler for /reverse
│   │   ├── file.go          // Handlers for /createfile, /deletefile
│   │   ├── status.go        // Handler for /status
│   │   └── simulate.go      // Handler for /simulate
│   ├── router/
│   │   └── router.go        // Routes HTTP requests to handlers
│   ├── server/
│   │   └── server.go        // TCP listener, reads/parses HTTP requests
│   └── utils/
│       └── recover.go       // RecoverAndRespond and other helpers
├── go.mod
└── README.md
```