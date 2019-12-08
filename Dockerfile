FROM golang:1.12-alpine

ENV HELM_LATEST_VERSION="v2.16.1"
ENV TERRATEST_LOG_PARSER_VERSION="v0.13.13"
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN apk add --update ca-certificates \
    && apk add --update -t deps wget git bzr openssl bash \
    && wget -q https://storage.googleapis.com/kubernetes-helm/helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
    && tar -xf helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
    && mv linux-amd64/helm /usr/local/bin \
    && rm /var/cache/apk/* \
    && rm -f /helm-${HELM_LATEST_VERSION}-linux-amd64.tar.gz \
    && wget -q -O /usr/local/bin/terratest_log_parser https://github.com/gruntwork-io/terratest/releases/download/${TERRATEST_LOG_PARSER_VERSION}/terratest_log_parser_linux_amd64 \
    && chmod +x /usr/local/bin/terratest_log_parser

WORKDIR /app/
COPY go.mod go.sum /app/
COPY test /app/test

RUN go test -i -tags=helm ./test

COPY docker-entrypoint.sh /usr/local/bin/

VOLUME /app/test

# TODO: Sort out a sane default for executing this

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
CMD ["./test/..."]

# CMD "/bin/sh" "-c" "go test -v -count=1 ./test/..."
