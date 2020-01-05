FROM golang:buster AS builder

WORKDIR /build
COPY . .
RUN apt update && apt install -y make
RUN make build

FROM debian:buster
WORKDIR /app
RUN mkdir /data
RUN apt update && apt install -y ca-certificates
COPY --from=builder /build/bin/mclogger /app

CMD ["./mclogger"]
