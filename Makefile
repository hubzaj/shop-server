IMG    := shop
LATEST := ${IMG}:latest

OWNER    := -OVERRIDE-WITH-ENV-OWNER-SUFFIX

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
	@helm template shop-service ./k8s/shop -f ./k8s/shop/values.yaml --set name=shop-service${OWNER},config.storage.host=postgres${OWNER} >> manifest/shop-service${OWNER}.yml

bake-on-demand-storage-manifest:
	@make create-manifest-dir
	@helm template postgres ./k8s/on-demand/storage --set name=postgres${OWNER} >> manifest/storage${OWNER}.yml

bake-on-demand-env-router-manifest:
	@make create-manifest-dir
	@helm template env-router ./k8s/on-demand/env-router >> manifest/env-router.yml

minikube-start:
	@minikube start
	@eval $(minikube docker-env)
	@minikube addons enable ingress

on-demand-deploy-env-router:
	@make bake-on-demand-env-router-manifest
	@kubectl apply -f manifest/env-router.yml

on-demand-deploy:
	@make bake-on-demand-storage-manifest
	@kubectl apply -f manifest/storage${OWNER}.yml
	@make bake-on-demand-shop-service-manifest
	@kubectl apply -f manifest/shop-service${OWNER}.yml

on-demand-env-router-url:
	@minikube service env-router --url

on-demand-env-router-teardown:
	@kubectl delete ingress env-router
	@kubectl delete service env-router
	@kubectl delete deployment env-router

on-demand-shop-service-teardown:
	@kubectl delete service shop-service${OWNER}
	@kubectl delete deployment shop-service${OWNER}
	@kubectl delete service postgres${OWNER}
	@kubectl delete statefulset postgres${OWNER}
	@kubectl delete configmap postgres${OWNER}-configuration

on-demand-teardown:
	@rm -rf manifest
	@make on-demand-env-router-teardown
	@make on-demand-shop-service-teardown

minikube-cleanup:
	@minikube delete
