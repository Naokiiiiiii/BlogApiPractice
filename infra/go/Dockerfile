FROM golang:1.23.3-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/BlogAppBackend

WORKDIR /go/src/github.com/BlogAppBackend

ADD ./ /go/src/github.com/BlogAppBackend

RUN go get -u golang.org/x/tools/cmd/goimports