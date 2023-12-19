IMG    := shop
LATEST := ${IMG}:latest

build:
	go build -o ./bin/shop ./cmd

run:
	@make build
	./bin/shop $(ARGS) start-shop-service

test-component:
	go test ./component-test/test

build-docker-image:
	@docker build -t ${LATEST} .

run-shop-in-docker:
	@docker run -e SHOP_HTTPSERVER_PORT=9000 -p 9000:9000 ${LATEST}
