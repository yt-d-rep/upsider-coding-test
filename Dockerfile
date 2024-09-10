FROM golang:1.23.0-bookworm

WORKDIR /root

RUN apt update && apt install -y git

WORKDIR /go/src/github.com/yt-d-rep/upsider-coding-test