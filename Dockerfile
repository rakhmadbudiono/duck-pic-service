FROM golang:1.17-alpine

WORKDIR /app
COPY . /app/

RUN go mod tidy
RUN go build -o app

CMD ["./app"]