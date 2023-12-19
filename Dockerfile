FROM golang:1.21

WORKDIR /usr/src/app

COPY . .

RUN go build -v -o /usr/local/bin/app ./go-top.go

CMD ["app"]

