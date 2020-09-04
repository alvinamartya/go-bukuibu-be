FROM golang:alpine

ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add --update -t curl go git
RUN go get -u -v github.com/gorilla/mux gorm.io/gorm github.com/joho/godotenv gorm.io/driver/postgres

CMD ["go", "run", "main.go"]