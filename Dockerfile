# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /go/src/github.com/dibrinsofor/urlplaylists
# COPY go.mod ./
# COPY go.sum ./
COPY . ./
RUN go mod download

ENV GO111MODULE=on
RUN go mod tidy

RUN go build -o /urlplaylists

EXPOSE 8080

CMD [ "/urlplaylists" ]
