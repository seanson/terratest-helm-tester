FROM golang:1.13-alpine

ENV HELM_LATEST_VERSION="v2.16.1"
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN apk add --update ca-certificates \
    && apk add --update -t deps wget git openssl bash \
    && wget -q https://storage.googleapis.com/kubernetes-helm/helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
    && tar -xf helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
    && mv linux-amd64/helm /usr/local/bin \
    && rm /var/cache/apk/* \
    && rm -f /helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz

WORKDIR /app/
COPY . /app

RUN go test -i -tags=helm ./test

VOLUME /app/test

# TODO: Sort out a sane default for executing this

# CMD "/bin/sh" "-c" "go test -v -count=1 ./test/..."
