FROM golang:latest

LABEL maintainer="Rafael Scheidt <rfscheidt@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build main.go

EXPOSE 8080

CMD ["./main"]