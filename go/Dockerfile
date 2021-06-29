FROM golang:1.15

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o test-wiismile .

EXPOSE 8080

ENTRYPOINT ["/app/test-wiismile"]