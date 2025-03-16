APP_NAME=client

.PHONY: build run clean

build:
	go build -o $(APP_NAME) .

manage: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME) manage

sync: build
	@./$(APP_NAME) sync

clean:
	rm -f $(APP_NAME)