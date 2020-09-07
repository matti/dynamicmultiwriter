FROM golang:1.15.0-alpine3.12
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /build
COPY . .
RUN go build
