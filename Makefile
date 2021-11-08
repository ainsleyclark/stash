# Setup
setup:
	go mod tidy && cd ./cmd go mod tidy
.PHONY: setup

# Build
build:
	cd ./cmd && go build -o luminati
.PHONY: build

# Runs main
run:
	cd ./cmd && go run main.go
.PHONY: run

# Run gofmt
format:
	go fmt ./...
.PHONY: format

# Run linter
lint:
	golangci-lint run ./
.PHONY: lint

# Test uses race and coverage
test:
	go clean -testcache && go test -race $$(go list ./... | grep -v /mocks/ | grep -v /cmd/) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

# Test with -v
test-v:
	go clean -testcache && go test -race -v $$(go list ./... | grep -v /mocks/ | grep -v /cmd/) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.out
.PHONY: cover

# Make mocks keeping directory tree
mock:
	rm -rf mocks && mockery --all --keeptree --exported=true
.PHONY: mock

# Make format, lint and test
all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
