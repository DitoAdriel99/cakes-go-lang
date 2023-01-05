
# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Dito <ditoadriel@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app

# Download necessary Go modules
COPY . .
EXPOSE 3030
RUN go mod tidy
RUN go build
CMD go run main.go