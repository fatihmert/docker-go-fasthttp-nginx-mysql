FROM golang:1.14.3-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

# install fasthhttp and fasthttp router
RUN go mod download

RUN go build -o main .

#EXPOSE 3000

CMD ["/app/main"]

# todo add healhtcheck