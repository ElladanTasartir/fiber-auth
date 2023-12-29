.PHONY: run-db
run-db:
	@docker-compose up -d

.PHONY: run
run: run-db
	@air -c ./build/.air.toml

.PHONY: build
build:
	@go build -o ./fiber-auth ./cmd/fiber-auth/main.go