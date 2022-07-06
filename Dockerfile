FROM golang:1.18.3-buster
WORKDIR /
COPY go.sum ./
COPY go.mod ./
RUN go mod download
COPY / .
CMD go run .