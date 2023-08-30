FROM golang:1.20.4

WORKDIR /usr/app

COPY . .

RUN go get -d -v ./...

RUN go build -o goGrocery .

EXPOSE 3000

CMD ["./goGrocery"]