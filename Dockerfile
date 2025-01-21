FROM golang:latest

WORKDIR /sshmachinebot

COPY . .

CMD [ "go","run","." ]
