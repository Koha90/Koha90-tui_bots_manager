build:
	@go build -o bin/bot-manager cmd/main.go

run: build
	@./bin/bot-manager

test:
	@go test ./...
