FROM golang:alpine

RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go mod download

CMD ["sh", "./start.sh"]

# goの開発をしたいときは永続化設定をして、こっちを使う
# RUN go install github.com/air-verse/air@latest
# CMD ["sh", "./dev.sh"]