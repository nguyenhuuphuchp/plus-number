# Makefile

APP_NAME=api
DB_HOST=localhost
DB_NAME=test
DB_SCHEMA=public

# Build binary
build:
	go build -o $(APP_NAME) ./cmd/api

# Run app
run: build
	./$(APP_NAME)

# Test bằng curl (chạy API GET /ping)
test:
	curl -v http://localhost:8080/ping

# Clean binary
clean:
	rm -f $(APP_NAME)
