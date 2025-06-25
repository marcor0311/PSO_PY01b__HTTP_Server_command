FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dispatcher ./cmd/dispatcher

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/dispatcher .
EXPOSE 8080                    
CMD ["./dispatcher"]
