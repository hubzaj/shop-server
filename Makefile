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

bake-shop-service-manifest:
	@helm template shop-service ./k8s/shop >> shop-service.yml

bake-storage-on-demand-manifest:
	@helm template postgres ./k8s/on-demand/storage >> storage.yml
