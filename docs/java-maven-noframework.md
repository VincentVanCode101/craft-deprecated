## Java Maven No Framework Template

## Overview
This document describes the files and structure of the `java-maven-noframework` template used by the `Craft` CLI tool. This template creates a lightweight Java project using Maven without any specific frameworks, enabling quick setup and easy customization. It also explains how to get the project up and running in a docker container

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
  Look for a container named `PROJECT_NAME-java-env`.

#### **2. Connect to the Development Container**
Once the container is running, connect to it for development purposes.

- **Open a bash session in the container:**
  ```bash
  docker exec -it PROJECT_NAME-java-env bash
  ```
  - use the `make` command from inside the container now ([see chapter](#build-and-run-with-makefile))

---

## Files Created

### **Directory Structure**
```
PROJECT_NAME/
├── pom.xml # Maven configuration file
├── src/
│   ├── main/
│   │   ├── java/
│   │   │   └── com/
│   │   │       └── main/
│   │   │           └── App.java # Main application entry point
│   │   └── resources/
│   │       └── application.properties # Configuration properties
│   └── test/
│       └── java/
│           └── com/
│               └── main/
│                   └── AppTest.java    # Sample test class
├── docker-compose.dev.yml  # Docker setup for local development
├── Makefile                # Build and run commands in Docker
├── README.md               # Initial project documentation
├── Dockerfile              # Dockerfile for building and running the application
├── .dockerignore           # Docker ignore file
└── .gitignore              # Git ignore file
```

### **File Descriptions**

#### 1. **pom.xml**
- The main Maven configuration file.
- Defines project dependencies, plugins, and build configurations.

#### 2. **src/main/java/com/main/App.java**
- The entry point of the application.
- Contains a basic `Hello World` program.

#### 3. **src/test/java/com/main/AppTest.java**
- A sample test class to verify the functionality of the application.

#### 4. **docker-compose.dev.yml**
- Docker Compose file for local development.
- Sets up a containerized environment for developing the application.

#### 5. **Makefile**
- Simplifies common tasks like building, running, testing, and cleaning up the project in a Dockerized environment.

#### 6. **README.md**
- Initial documentation for the generated project.
- Explains how to use the Makefile and other features.

#### 7. **Dockerfile**
- Builds a Docker image for the Java application.
- Sets up a containerized environment for running the application.

#### 8. **.dockerignore**
- Specifies files and directories to exclude from the Docker build context.
- Optimizes Docker image builds by ignoring unnecessary files.

#### 9. **.gitignore**
- Specifies intentionally untracked files to ignore in the Git repository.
- Common entries include build artifacts and dependency directories.

---

## Build and Run with Makefile

The generated project includes a `Makefile` that simplifies the following operations:
- can only be used from *inside the container*

- **Compile the project**:
  ```bash
  make
  ```
- **Build specific Java files**:
  ```bash
  make build ARGS=src/main/java/com/main/foo/bar.java
  ```
- **Run the application**:
  ```bash
  make run
  ```
- **Create an executable JAR**:
  ```bash
  make uber-jar
  ```
- **Run the JAR**:
  ```bash
  make run-jar
  ```
- **Execute all tests**:
  ```bash
  make test
  ```
- **Clean up build artifacts**:
  ```bash
  make clean
  ```

For full documentation of the Makefile, refer to the `README.md` in the generated project.

---

## Notes
- The default main class is set to `com.main.App`. Update the `MAIN_CLASS` variable in the Makefile if your main class differs.
- The project name is automatically set during initialization. Update the `JAR_NAME` variable in the Makefile if you want to change the generated JAR's name.

---
When using the Craft tool, the specified name (when creating a new project) is propagated and used consistently throughout the entire application.

With this template, you can quickly set up a Java Maven project and begin development in a containerized environment with minimal configuration.
