FROM golang:1.23.0-bookworm

WORKDIR /root

RUN apt update && apt install -y git && \
  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-arm64.tar.gz | tar xvz && \
  mv migrate /usr/local/bin/

WORKDIR /go/src/github.com/yt-d-rep/upsider-coding-test