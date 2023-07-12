FROM golang:alpine3.18 AS builder

COPY . /QuizBot/

WORKDIR /QuizBot/

RUN go mod download

RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 QuizBot/bin/bot .
COPY --from=0 QuizBot/configs configs/

EXPOSE 80

CMD ["./bot"]