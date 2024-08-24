FROM golang:1.22.1-alpine3.19
LABEL vendor=Menoudo\ Software \
      com.example.is-production="Yes" \
      com.example.version="1.00.00" \
      com.example.release-date="2024-03-09"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY todo-tg-bot /go/todo-tg-bot/todo-tg-bot .
CMD [ "/app/todo-tg-bot" ]
