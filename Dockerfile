FROM golang:alpine

ADD . /go/src/go-bukuibu-be
WORKDIR /go/src/go-bukuibu-be

RUN apk add --update -t curl go git
RUN go get -u -v github.com/gorilla/mux gorm.io/gorm github.com/joho/godotenv gorm.io/driver/postgres

CMD ["go", "run", "main.go"]