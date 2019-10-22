FROM golang:1.11.4

WORKDIR /go/src/github.com/pugchat_multimedia
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh

EXPOSE 8081
