FROM golang:alpine

RUN go get -d -v ./...

WORKDIR /go/src/app
COPY . .

CMD ["go", "run", "main.go"]