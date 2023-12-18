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
	@docker build -f ./docker/Dockerfile -t ${LATEST} .

run-shop-in-docker:
	@docker run ${LATEST} start-shop-service
