FROM golang:1.9-alpine

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git \
    && rm -rf /var/cache/apk/* \
    && go get github.com/ssimunic/gossm/cmd/gossm \
    && go build github.com/ssimunic/gossm/cmd/gossm \
    && mv gossm /usr/local/bin/ \
    && mkdir -p /configs /var/log/gossm

ADD configs /configs

CMD ["gossm", "-config", "/configs/default.json", "-http", ":8080", "-log", "/var/log/gossm/gossm.log"]

EXPOSE 8080
