FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY config ./config
COPY cmd ./cmd
COPY pkg ./pkg
RUN go build -o ./bin/shop ./cmd

CMD ["bin/shop", "start-shop-service"]
