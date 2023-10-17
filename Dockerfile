FROM golang:latest

ENV GO111MODULE=on
ENV PORT=8080
WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download
RUN go build -o main .
RUN chmod +x main
