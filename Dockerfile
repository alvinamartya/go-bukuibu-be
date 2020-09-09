# Start from golan base image
FROM golang:alpine as builder

ENV GO111MODULE=auto

# Add Maintainer info
LABEL maintainer="Alvin Amartya <alvinamartya1@gmail.com>"

# Install git
# Git is required for fetching the dependencies
RUN apk update && apk add --no-cache git go bash

# Set the current working directory inside the container
RUN mkdir /app
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/db.sql .
COPY --from=builder /app/config ./config

# Expose port 8080 to the outside WORD
EXPOSE 8080

# Command to run the executable
CMD ["./main"]