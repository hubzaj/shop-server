build:
	go build -o ./bin/shop ./cmd

run:
	@make build
	./bin/shop $(ARGS) shop

test-component:
	go test ./component-test/test
