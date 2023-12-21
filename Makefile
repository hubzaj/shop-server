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

bake-on-demand-shop-service-manifest:
	@helm template shop-service ./k8s/shop -f ./k8s/shop/values-on-demand.yaml>> shop-service.yml

bake-on-demand-storage-manifest:
	@helm template postgres ./k8s/on-demand/storage >> storage.yml

minikube-start:
	@minikube start

deploy-on-demand:
	@make bake-on-demand-storage-manifest
	@kubectl apply -f storage.yml
	@make bake-on-demand-shop-service-manifest
	@kubectl apply -f shop-service.yml

expose-shop-service-url:
	@minikube service shop-service --url

cleanup:
	@rm shop-service.yml
	@kubectl delete deployment shop-service
	@kubectl delete service shop-service
	@rm storage.yml
	@kubectl delete statefulset postgres
	@kubectl delete service postgres
	@kubectl delete configmap postgres-configuration

minikube-cleanup:
	@minikube delete
