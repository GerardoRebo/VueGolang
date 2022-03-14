FROM golang:1.17.8-alpine3.15

WORKDIR /golang-docker

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./golang-docker"