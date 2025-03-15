APP_NAME=server

.PHONY: build run clean

build:
	go build -o $(APP_NAME) .

run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)

clean:
	rm -f $(APP_NAME)