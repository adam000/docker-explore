FROM golang:1.15.3

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

CMD ["./app"]
