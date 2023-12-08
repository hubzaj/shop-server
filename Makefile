build:
	go build -o ./bin/shop ./cmd

run:
	@make build
	./bin/shop $(ARGS) shop

test-component:
	gotestsum --format short-verbose ./component-test/... -parallel=64
