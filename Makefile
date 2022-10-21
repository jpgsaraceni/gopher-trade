.PHONY: lint
lint:
	@echo "==> installing golangci-lint..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.50.0
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners -c ./.golangci.yml ./...

.PHONY: test
test:
	@echo "==> running go tests..."
	go test -v -race ./...

.PHONY: install-tools
install-tools:
	@echo "==> installing development tools..."
	go install github.com/matryer/moq@v0.2.7
	go install github.com/swaggo/swag/cmd/swag@v1.8.4

.PHONY: clean-generated
clean-generated:
	@echo "==> removing old generated files..."
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +
	rm -rf gen/*

.PHONY: generate
generate: clean-generated install-tools
	@echo "==> tidying go mod..."
	go mod tidy
	@echo "==> running go generate..."
	go generate ./...
	@echo "==> generating swagger..."
	swag init -g /cmd/api/main.go

.PHONY: api
api:
	@echo "==> running Gopher Trade API on docker container..."
	docker-compose up -d --build

.PHONY: stop
stop:
	@echo "==> stopping Gopher Trade API..."
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs app

.PHONY: all-logs
all-logs:
	docker-compose logs

.PHONY: live-logs
live-logs:
	docker-compose logs -f app

.PHONY: load-test
load-test:
	@echo "==> downloading load test utility..."
	go install -v go.ddosify.com/ddosify@latest
	ddosify --config load_test/config.json