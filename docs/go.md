## Go Template

## Overview
This document describes the files and structure of the Go template provided by the `Craft` CLI tool. This template sets up a Go project with essential configuration files, Docker support, and a Makefile for streamlined development. It also explains how to get the project up and running in a docker container

---

### **How to Start the Project Using Docker**

This project is configured to run inside a Docker container for consistent development environments. Follow the steps below to set up, start, and use the project.

---

### **Steps to Start the Project**

#### **1. Build and Start the Docker Environment**
Use the provided `docker-compose.dev.yml` file to build and start the development container.

- **Build the container:**
  ```bash
  docker-compose -f docker-compose.dev.yml build
  ```

- **Start the container:**
  ```bash
  docker-compose -f docker-compose.dev.yml up -d
  ```

- **Confirm the container is running:**
  ```bash
  docker ps
  ```
  Look for a container named `PROJECT_NAME-go-compiler`.

#### **2. Connect to the Development Container**
Once the container is running, connect to it for development purposes.

- **Open a bash session in the container:**
  ```bash
  docker exec -it PROJECT_NAME-go-compiler bash
  ```

---

## Files Created

### **Directory Structure**
```
PROJECT_NAME/
├── .gitignore              # Git ignore file
├── .dockerignore           # Docker ignore file
├── docker-compose.dev.yml  # Docker Compose configuration for local development
├── Dockerfile              # Dockerfile for building and running the application
├── go.mod                  # Go module configuration file
├── go.sum                  # Go module checksum file
├── main.go                 # Main entry point for the Go application
├── Makefile                # Build and run commands for the Go project
└── pre-commit              # Pre-commit hook for code quality checks
```

### **File Descriptions**

#### 1. **.gitignore**
- Specifies intentionally untracked files to ignore in the Git repository.
- Common entries include build artifacts and dependency directories.

#### 2. **.dockerignore**
- Specifies files and directories to exclude from the Docker build context.
- Optimizes Docker image builds by ignoring unnecessary files.

#### 3. **docker-compose.dev.yml**
- Configures a Dockerized development environment for the Go application.
- Defines services and volumes for containerized development.

#### 4. **Dockerfile**
- Builds a Docker image for the Go application.
- Sets up a containerized environment for running the application.

#### 5. **go.mod**
- The `go.mod` file defines the module name and dependencies.
- The module name is dynamically replaced with the project name during initialization.

#### 6. **go.sum**
- The `go.sum` file provides checksums for module dependencies.

#### 7. **main.go**
- The main entry point of the Go application.
- Includes a simple `Hello, World!` program as a starting point.

#### 8. **Makefile**
- Simplifies common operations for Go projects, such as building, running, and testing the application.

#### 9. **pre-commit**
- A pre-commit hook to enforce code quality and formatting checks.
- Can be linked into the Git hooks directory for automatic execution before commits.

---

## Using the Makefile

The generated project includes a `Makefile` to streamline development tasks:

- **Build the application**:
  ```bash
  make build
  ```
- **Run the application**:
  ```bash
  make run
  ```
- **Run tests**:
  ```bash
  make test
  ```
- **Clean build artifacts**:
  ```bash
  make clean
  ```

For full documentation of the Makefile, refer to the `README.md` in the generated project.

---

## Notes
- The `go.mod` and `go.sum` files are customized during initialization to include the project name.
- The `main.go` file serves as the starting point for your application; modify it as needed to suit your requirements.

---


With this template, you can quickly set up and manage a Go project with minimal configuration effort.
