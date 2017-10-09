FROM alpine:edge

ENV GOPATH /go
ENV PATH /go/src/github.com/ssimunic/gossm/bin:$PATH

ADD . /go/src/github.com/ssimunic/gossm

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
    && apk add --no-cache --update bash ca-certificates \
    && apk add --no-cache --virtual .build-deps go gcc git libc-dev \
    && mkdir -p /configs /usr/local/bin /var/log/gossm \
    && go get github.com/gregdel/pushover \
    && cd /go/src/github.com/ssimunic/gossm \
    && go build -v -o /usr/local/bin/gossm cmd/gossm/main.go \
    && apk del --purge .build-deps \
    && rm -rf /var/cache/apk*

ADD configs /configs

CMD ["gossm", "-config", "/configs/default.json", "-http", ":8080", "-log", "/var/log/gossm/gossm.log"]

EXPOSE 8080
