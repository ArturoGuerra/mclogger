FROM golang:buster AS builder

WORKDIR /build
COPY . .
RUN apk add --update make
RUN make build

FROM debian:buster
WORKDIR /app
COPY --from=builder /build/mclogger /app

CMD ["./mclogger"]
