build:
	@go build -o bin/ecom main.go

test:
	@go test -v ./...

run:
	@./bin/ecom

migration:
	@migration create -ext sql -dir cmd/migrate/migrations $filter-out $@, $(MAKECMDGOALS)

migrate-up:
	@go run cmd/migrate/main.go

migrate-down:
	@go run cmd/migrate/main.go down