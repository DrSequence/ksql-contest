.PHONY: build run-local

APP_NAME := order-producer
BUILD_PATH := ./bin

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_PATH)/$(APP_NAME) .

run-local: build
	@echo "Running $(APP_NAME) locally..."
	@$(BUILD_PATH)/$(APP_NAME)
