BINARY_NAME := crafter
MAIN_PACKAGE := ./main.go

.PHONY: all build run clean

all: build

build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	@echo "Running the project..."
	./$(BINARY_NAME) $(ARGS)

clean:
	go clean
