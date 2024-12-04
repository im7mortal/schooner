FROM registry.hub.docker.com/library/golang:1.23.4-bookworm as build

WORKDIR /build

COPY . .

RUN go test