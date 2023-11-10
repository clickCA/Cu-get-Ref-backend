FROM golang:1.21.0-alpine as builder

WORKDIR /src

COPY . .

EXPOSE 8080

RUN go mod tidy

CMD [ "go","run","main.go" ]