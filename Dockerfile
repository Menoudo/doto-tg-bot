FROM golang:1.22.1-alpine3.19
EXPOSE 8080
RUN apk add --update git; \
    apk add curl; \
    mkdir -p ${GOPATH}/todo-tg-bot
WORKDIR ${GOPATH}/todo-tg-bot/
COPY test-rest-api.go ${GOPATH}/todo-tg-bot/
COPY go.* ${GOPATH}/todo-tg-bot/
RUN cd ${GOPATH}/todo-tg-bot/ && \
    go build -o todo-tg-bot .

# Linter
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3
RUN golangci-lint --version
RUN golangci-lint run

#
FROM golang:1.22.1-alpine3.19
LABEL vendor=Menoudo\ Software \
      com.example.is-production="Yes" \
      com.example.version="1.00.00" \
      com.example.release-date="2024-03-09"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/todo-tg-bot/todo-tg-bot .
CMD [ "/app/todo-tg-bot" ]