FROM golang:1.10.2-alpine3.7

RUN apk add --no-cache curl wget bash git

# install dep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# install retool
WORKDIR /tmp
RUN wget -q https://github.com/twitchtv/retool/archive/v1.3.7.tar.gz
RUN mkdir -p /go/src/github.com/twitchtv/retool && tar zxf v1.3.7.tar.gz -C /go/src/github.com/twitchtv/retool --strip-components=1
WORKDIR /go/src/github.com/twitchtv/retool
RUN go install

WORKDIR /go/src/app

# ensure dependencies are up to date
COPY ./Gopkg.toml ./Gopkg.lock ./
RUN dep ensure -vendor-only

# ensure tools are up to date
COPY ./tools.json ./
RUN retool sync

COPY . .
