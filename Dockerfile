FROM golang:1.11-alpine

RUN set -xe \
    && apk add --no-cache postgresql-client mariadb-client git build-base

WORKDIR /go/src/github.com/tsuty/table-doc

VOLUME /go/src/github.com/tsuty/table-doc

