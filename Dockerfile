FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./go.mod
COPY lib ./lib
COPY merrymake.go ./merrymake.go

RUN go mod download

RUN go build
