FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
RUN go get -u github.com/amacneil/dbmate

FROM alpine:3.9
WORKDIR /go/src/github.com/snowman-mh/go-sample
COPY --from=builder /go/bin/dbmate /usr/local/bin/dbmate
COPY db /go/src/github.com/snowman-mh/go-sample
