FROM golang:alpine

WORKDIR /go/src/app
COPY images .

RUN go get -d -v ./...
RUN go install -v ./...

CMD["go", "run", "main.go"]