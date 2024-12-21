# Makefile for Quarkus Project

# Default target
.DEFAULT_GOAL := help

# Variables
MVNW = mvnw

# Help target
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application (creates an executable JAR)"
	@echo "  make dev          - Run the application in development mode"
	@echo "  make test         - Run tests"
	@echo "  make package      - Package the application"

# Build the application
build:
	$(MVNW) package

# Run the application in development mode
dev:
	$(MVNW) quarkus:dev

# Run tests
test:
	$(MVNW) test

# Package the application
package:
	$(MVNW) package -Dquarkus.package.jar.type=uber-jar
