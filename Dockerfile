FROM golang:rc-alpine AS builder
RUN mkdir -p /go/src/github.com/vlcty/example
WORKDIR src/github.com/vlcty/example
COPY main.go .
RUN go build -o color main.go
ENTRYPOINT ["/go/src/github.com/vlcty/example/color"]