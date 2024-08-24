FROM golang:1.22.1-alpine3.19
EXPOSE 8080
RUN apk add --update git; \
    apk add curl; \
    mkdir -p ${GOPATH}/todo-tg-bot
WORKDIR ${GOPATH}/todo-tg-bot/
COPY *.go ${GOPATH}/todo-tg-bot/
COPY go.* ${GOPATH}/todo-tg-bot/
RUN cd ${GOPATH}/todo-tg-bot/ && \
    go build -o todo-tg-bot .

# Linter
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3
