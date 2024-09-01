FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git build-base

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache sqlite

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN adduser -D appuser

RUN chown -R appuser:appuser /app

USER appuser

RUN which air

CMD ["air", "-c", ".air.toml"]