###################
# Linting         #
###################
.PHONY: fmt
fmt:
	golangci-lint fmt ./...

.PHONY: lint
lint: fmt
	golangci-lint run
