FROM golang:1.18.0-alpine3.15
WORKDIR /
COPY / .
CMD go run .