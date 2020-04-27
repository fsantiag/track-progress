FROM golang:1.14.2

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir /go/src/app

COPY . /go/src/app

WORKDIR /go/src/app

RUN dep ensure

RUN go build -o /app/main .

CMD ["/app/main"]
