# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /app

COPY go.build.mod ./go.mod
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /urlplaylists

EXPOSE 8080

CMD [ "/urlplaylists" ]
