FROM golang:1.20 AS builder
WORKDIR /src
COPY . .
RUN set -e; \
    go mod tidy; \
    go build -o ./bin/pag .

FROM debian:11
RUN set -eux; \
    apt-get update; \
    apt-get install -y apt-transport-https ca-certificates
COPY --from=builder /src/bin/pag /usr/bin/pag
EXPOSE 9000
ENTRYPOINT ["pag"]