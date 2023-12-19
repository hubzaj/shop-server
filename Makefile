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

start-container-development-environment:
	@docker-compose up

stop-container-development-environment:
	@docker-compose down
