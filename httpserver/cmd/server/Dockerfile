FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o httpserver ./cmd/server

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/httpserver .
EXPOSE 8080                    
CMD ["./httpserver"]
