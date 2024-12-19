---
# PROJECT_NAME
### **How to Start the Project Using Docker**

This project is configured to run inside a Docker container for consistent development environments. Follow the steps below to set up, start, and use the project.

---

### **Steps to Start the Project**

#### **1. Build and Start the Docker Environment**
Use the provided `docker-compose.dev.yml` file to build and start the development container.

- **Build the container:**
  ```bash
  docker compose -f docker-compose.dev.yml build
  ```

- **Start the container:**
  ```bash
  docker compose -f docker-compose.dev.yml up -d
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

### **How to Use the Makefile (Container Usage)**

This `Makefile` is designed to simplify building, running, and cleaning up a Go application. It includes commands for building the application for local and Linux environments, running the binary, and cleaning up build artifacts.

---

### **Commands Overview**

#### **1. Default Target: `make` or `make all`**
- **Purpose**: Builds the main Go project binary.
- **Usage**:
  ```bash
  make
  ```
- **Effect**:
  - Compiles the main Go application specified by `MAIN_PACKAGE` (`./main.go`) into a binary named `PROJECT_NAME`.

---

#### **2. Build: `make build`**
- **Purpose**: Builds the Go application or a specified Go file.
- **Usage**:
  - **Build the main project**:
    ```bash
    make build
    ```
  - **Build a specific Go file**:
    ```bash
    make build ARGS=path/to/otherfile.go
    ```
- **Effect**:
  - If no `ARGS` is provided, compiles `MAIN_PACKAGE` into the binary `$(BINARY_NAME)`.
  - If `ARGS` is provided, compiles the specified file into a binary with the same name (without the `.go` extension).

- **Example**:
  ```bash
  make build
  ```
  Output:
  ```
  Building the main project (./main.go)...
  ```

---

#### **3. Build for Linux: `make linux-build`**
- **Purpose**: Builds the Go application for a Linux environment.
- **Usage**:
  ```bash
  make linux-build
  ```
- **Effect**:
  - Sets environment variables (`CGO_ENABLED=0` and `GOOS=linux`) for a Linux-compatible build.
  - Compiles the `MAIN_PACKAGE` into the binary `$(BINARY_NAME)`.

- **Example**:
  ```bash
  make linux-build
  ```
  Output:
  ```
  Building for Linux (CGO_ENABLED=0 GOOS=linux)...
  ```

---

#### **4. Run: `make run`**
- **Purpose**: Runs the compiled Go binary or a specified binary.
- **Usage**:
  - **Run the main binary**:
    ```bash
    make run
    ```
  - **Run a specific binary**:
    ```bash
    make run ARGS=path/to/otherfile.go
    ```
- **Effect**:
  - Executes the `$(BINARY_NAME)` binary if `ARGS` is not provided.
  - If `ARGS` is provided, runs the binary corresponding to the specified Go file.

- **Example**:
  ```bash
  make run
  ```
  Output:
  ```
  Running the main project (PROJECT_NAME)...
  ```

---

#### **5. Clean: `make clean`**
- **Purpose**: Cleans up build artifacts.
- **Usage**:
  ```bash
  make clean
  ```
- **Effect**:
  - Executes `go clean` to remove any intermediate or build artifacts created during the build process.

- **Example**:
  ```bash
  make clean
  ```
  Output:
  ```
  Cleaning up Go build artifacts...
  ```

---

### **Best Practices**
- **Binary Name**: Update the `BINARY_NAME` variable to reflect your application name.
- **Main Package**: Ensure the `MAIN_PACKAGE` points to your main Go file (default is `./main.go`).

This `Makefile` simplifies project management by providing quick commands for building, running, and cleaning your Go application, as well as preparing it for Linux deployment.