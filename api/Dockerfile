FROM golang:1.13.4-alpine3.10 AS build

WORKDIR /
COPY . .

RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.10/main > /etc/apk/repositories; \
    echo http://mirror.yandex.ru/mirrors/alpine/v3.10/community >> /etc/apk/repositories

RUN apk update \
    && apk add --no-cache git \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/google/uuid \
    && go get github.com/gin-gonic/gin
RUN go build -o bin/main main.go

FROM alpine:3.10
COPY --from=build bin/main /usr/local/bin
CMD ["main"]
