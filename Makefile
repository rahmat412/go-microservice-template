build-img:
	@echo $(GIT_TOKEN) | docker build --secret id=github_token,src=/dev/stdin -t go-microservice-template .
run-stack:
	@docker compose up -d

stop-stack:
	@docker compose down

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

sync-jet:
	jet -dsn="postgresql://postgres:StrongPassword123@localhost:5432/go-microservice-template?sslmode=disable" -schema=public -path=./database/

new-migration:
	@read -p "Enter migration name: " name; \
	goose -dir ./database/migrations create $$name sql

go-gen:
	@go generate ./...

build:
	@go build -o service_template ./cmd

run:
	@./service_template

run-api:
	@go run ./cmd/main.go