MIGRATIONS_PATH="db/migration"
DATABASE_URL=postgres://postgres:postgres@localhost:5432/pack?sslmode=disable

###################
# Database        #
###################
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database "$(DATABASE_URL)" down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

###################
# Testing         #
###################
.PHONY: mock
mock:
	mockery --output test/gen/example/mocks --dir internal/example --all\

.PHONY: test
test: mock
	go test ./...

###################
# Linting         #
###################
.PHONY: fmt
fmt:
	golangci-lint fmt ./...

.PHONY: lint
lint: fmt
	golangci-lint run
