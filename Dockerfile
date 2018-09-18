#build
FROM golang:1.10 AS build

WORKDIR /go/src/github.com/jonnywei/multistage-go

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

CMD ["./multistage-go"]