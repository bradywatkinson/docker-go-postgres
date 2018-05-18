FROM golang:1.10.2-alpine3.7

RUN apk add --no-cache curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN apk add --no-cache bash git

RUN go get -u golang.org/x/lint/golint

WORKDIR /go/src/app
ADD ./app /go/src/app

RUN dep ensure
