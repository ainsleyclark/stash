# Setup
setup:
	go mod tidy && cd ./cmd go mod tidy
.PHONY: setup

# Redis
example-redis:
	cd ./cmd && go run main.go --redis
.PHONY: example-redis

# Memory
example-memory:
	cd ./cmd && go run main.go --memory
.PHONY: example-memory

# Memcache
example-memcache:
	cd ./cmd && go run main.go --memcache
.PHONY: example-memory

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
	rm -rf mocks && mockery --all --dir="./test"
.PHONY: mock

# Make format, lint and test
all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
