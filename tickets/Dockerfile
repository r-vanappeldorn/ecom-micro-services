FROM golang:1.18-buster

WORKDIR /app

COPY . .

RUN go build -o cmd/tickets

EXPOSE 8080

CMD ["./cmd/tickets"]