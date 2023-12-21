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

create-manifest-dir:
	[ -d "manifest" ] || mkdir -p "manifest"

bake-shop-service-manifest:
	@make create-manifest-dir
	@helm template shop-service ./k8s/shop >> manifest/shop-service.yml

bake-on-demand-shop-service-manifest:
	@make create-manifest-dir
	@helm template shop-service ./k8s/shop -f ./k8s/shop/values-on-demand.yaml>> manifest/shop-service.yml

bake-on-demand-storage-manifest:
	@make create-manifest-dir
	@helm template postgres ./k8s/on-demand/storage >> manifest/storage.yml

minikube-start:
	@minikube start

deploy-on-demand:
	@make bake-on-demand-storage-manifest
	@kubectl apply -f manifest/storage.yml
	@make bake-on-demand-shop-service-manifest
	@kubectl apply -f manifest/shop-service.yml

expose-shop-service-url:
	@minikube service shop-service --url

on-demand-cleanup:
	@rm -rf manifest
	@kubectl delete deployment shop-service
	@kubectl delete service shop-service
	@kubectl delete statefulset postgres
	@kubectl delete service postgres
	@kubectl delete configmap postgres-configuration

minikube-cleanup:
	@minikube delete
