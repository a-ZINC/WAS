build:
	@go build -o bin/was
run-agent: build
	@./bin/was agent
run-broker: build
	@./bin/was broker
test:
	@go test ./...