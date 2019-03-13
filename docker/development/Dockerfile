FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/github.com/snowman-mh/go-sample

ENV GO111MODULE on

RUN apk add --no-cache alpine-sdk git

RUN go get -u github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
