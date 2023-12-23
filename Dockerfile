FROM golang:1.21.0

WORKDIR /go/src/app
COPY . /go/src/app
RUN go build -o main .

ENTRYPOINT ["/app/main"]