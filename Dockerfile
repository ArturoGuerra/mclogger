FROM golang:buster AS builder

WORKDIR /build
COPY . .
RUN apt update && apt install build-essential
RUN make build

FROM debian:buster
WORKDIR /app
COPY --from=builder /build/mclogger /app

CMD ["./mclogger"]
