FROM golang:1.16-alpine

WORKDIR /app

ADD .  /app

RUN go build -o /terminal

EXPOSE 9090

CMD [ "/terminal" ]