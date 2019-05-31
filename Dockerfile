FROM golang:alpine AS development
WORKDIR $GOPATH/src
COPY . .
RUN go build -o app