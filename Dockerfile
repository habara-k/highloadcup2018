FROM golang:latest

RUN apt-get update
RUN apt-get install -y redis-server supervisor

RUN go get github.com/go-redis/redis

ADD main.go /go/src/github.com/habara-k/highloadcup2018/main.go

WORKDIR /go/src/github.com/habara-k/highloadcup2018

EXPOSE 80

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
CMD ["/usr/bin/supervisord"]
