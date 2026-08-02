FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /auth-gateway .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

COPY --from=builder /auth-gateway /auth-gateway

ENV HTTP_ADDR=:8080
ENV UPSTREAM=http://todoapp:5050

EXPOSE 8080

ENTRYPOINT ["/auth-gateway"]
