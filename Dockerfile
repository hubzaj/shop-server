FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o ./bin/shop ./cmd

EXPOSE 8080

CMD ["/bin/shop", "shop"]