FROM golang:1.25.1-alpine AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  APP_NAME=cmit

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /${APP_NAME} ./cmd/cmit

FROM alpine:latest

RUN apk add --no-cache git bash

COPY --from=builder /cmit /usr/local/bin/cmit

ENTRYPOINT ["cmit"]
