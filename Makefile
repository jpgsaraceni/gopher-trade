.PHONY: lint
lint:
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.44.2
endif
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners -c ./.golangci.yml ./...

.PHONY: test
test:
	@echo "Running go tests"
	go test -v ./...

.PHONY: install-tools
install-tools:
	@echo "Installing development tools"
	go install github.com/matryer/moq@latest

.PHONY: clean-generated
clean-generated:
	@echo "Removing old generated files"
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +
	rm -rf gen/*

.PHONY: generate
generate: clean-generated install-tools
	@echo "Tidying go mod"
	go mod tidy
	@echo "Running go generate"
	go generate ./...