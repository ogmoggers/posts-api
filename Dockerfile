FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/posts-api ./cmd/api

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate && \
    apk del curl

COPY --from=builder /app/posts-api .
COPY --from=builder /app/schema ./schema
COPY --from=builder /app/configs ./configs

RUN adduser -D -g '' appuser
USER appuser

CMD ["./posts-api"]
