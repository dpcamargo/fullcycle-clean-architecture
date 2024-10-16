FROM golang:1.18 as builder
   
WORKDIR /app
COPY . .

RUN go get -u github.com/pressly/goose/v3

CMD ["goose", "up"]