FROM golang:buster AS builder

WORKDIR /build
COPY . .
RUN go build -o bin/mclogger main.go

FROM debian:buster
WORKDIR /app
COPY --from=builder /build/bin/mclogger /app

CMD ["./mclogger"]
