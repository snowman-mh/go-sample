FROM golang:1.12.0 AS builder
WORKDIR /go/src/github.com/snowman-mh/go-sample
ENV GO111MODULE on
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-sample

FROM scratch
COPY --from=builder /go/bin/go-sample /go/bin/go-sample
EXPOSE 8080
ENTRYPOINT ["/go/bin/go-sample"]
