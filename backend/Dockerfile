#
# Build the Go server
#
FROM golang:1.13-alpine3.11 as go-build

WORKDIR /go/githubstars/

RUN apk update && apk add git gcc musl-dev

COPY app/*.go ./
COPY app/go.mod ./
COPY app/go.sum ./

ENV GOOS=linux
ENV GO111MODULE=on

RUN go build

EXPOSE 8000

VOLUME [ "/go/githubstars/data/" ]

ENTRYPOINT [ "githubstars" ]
