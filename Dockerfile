FROM golang:1.13.3

ADD . /go/src/github.com/shivakar/http-test-server
WORKDIR /go/src/github.com/shivakar/http-test-server

RUN go install -mod vendor -x github.com/shivakar/http-test-server/cmd/...

CMD server
