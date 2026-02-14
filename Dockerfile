# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/api-server ./cmd/api-server

FROM alpine:3.22
WORKDIR /app
RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /out/api-server /app/api-server

ENV APP_ENV=production
ENV HTTP_ADDR=:8080
EXPOSE 8080

USER app
ENTRYPOINT ["/app/api-server"]
