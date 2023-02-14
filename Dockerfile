FROM golang:1.18

WORKDIR /usr/src/app
COPY . .

RUN go get .
RUN go build .

VOLUME [ "/data" ]

CMD ["/usr/src/app/main"]
