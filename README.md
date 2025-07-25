# User Manual – HTTP Server (RedUnix)

## Requirements
- Docker
- Git
- Postman (optional, for testing)

## Installation

```bash
git clone https://github.com/marcor0311/PSO_PY01b__HTTP_Server_command.git
```

## Running the Server

Build Dispatcher and Worker images
```bash
docker compose build
```

Start Dispatcher (exposed on localhost:8080) + 3 Worker replicas
```bash
docker compose up -d
```

Logs for all workers
```bash
docker-compose logs --follow dispatcher  
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


## Parallel Problems

### Monte Carlo 
- **Endpoint:** `/montecarlo?points=N`
- **Description:** Estimates π by distributing N random point simulations across workers.
- **Command:**

```bash
  # replace 1000000 with the total number of random points to use
  curl "http://localhost:8080/montecarlo?points=1000000"
```

### Word Count
- **Endpoint:** `/countwords?url=<FILE_URL>`
- **Description:** Counts the frequency of words in a large file by distributing chunks across workers.
- **Command:**

```bash
  # replace <FILE_URL> with the URL of your large file
  curl "http://localhost:8080/countwords?url=<FILE_URL>"
```

```bash
  # Example: Linux device documentation in plain text
  curl "http://localhost:8080/countwords?url=https://www.kernel.org/doc/Documentation/admin-guide/devices.txt"
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
├── cmd/
│   ├── dispatcher/
│   │   └── main.go
│   └── server/
│       └── main.go
├── internal/
│   ├── constants/
│   │   └── paths.go
│   ├── dispatcher/
│   │   └── dispatcher.go
│   ├── handlers/
│   │   └── files.go
│   │   └── files_test.go
│   │   └── help.go
│   │   └── math.go
│   │   └── math_test.go
│   │   └── simulate.go
│   │   └── simulate_test.go
│   │   └── strings.go
│   │   └── strings_test.go
│   │   └── system.go
│   │   └── system_test.go
│   │   └── worker.go
│   │   └── worker_test.go
│   ├── router/
│   │   └── router.go
│   │   └── routerDispatcher.go
│   │   └── routerParallel.go
│   │   └── routerServer.go
│   ├── tcp/
│   │   └── connection.go
│   │   └── client.go
│   ├── utils/
│   │   └── http.go
│   │   └── http_test.go
│   │   └── helpers.go
│   │   └── helpers_test.go
│   └── worker/
│       └── worker.go
├── postman/
│   └── server.postman_collection.json
├── go.mod
├── docker-compose.yml
```
