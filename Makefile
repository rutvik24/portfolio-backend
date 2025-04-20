# Makefile for building and running the Go backend project

# Variables
APP_NAME := portfolio-backend
CMD_DIR := cmd
MAIN_FILE := $(CMD_DIR)/main.go

# Targets
.PHONY: run build clean

run:
	@echo "Running the project..."
	go run $(MAIN_FILE)

build:
	@echo "Building the project..."
	go build -o $(APP_NAME) $(MAIN_FILE)

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)