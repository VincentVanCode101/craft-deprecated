# {PROJECT_NAME}
### **How to Start the Project Using Docker**
This project is configured to run inside a Docker container for consistent development environments. Follow the steps below to set up, start, and use the project.

---

### **Steps to Start the Project**

#### **1. Build and Start the Docker Environment**
Use the provided `docker-compose.dev.yml` file to create and start the development container.

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
  Look for a container named `{PROJECT_NAME}-rust-env`.

#### **2. Connect to the Development Container**
Once the container is running, connect to it for further development.

- **Open a bash session in the container:**
  ```bash
  docker exec -it {PROJECT_NAME}-rust-env bash
  ```

#### **3. Use the Makefile for Project Operations**
After connecting to the container, you can use the `Makefile` to build, run, and test the application.


### **How to Use the Makefile (Container Usage)**

This `Makefile` is designed to streamline the process of building, running, testing, and cleaning up a Rust project inside a Docker container environment.

You need to connect to the [development container](#2-connect-to-the-development-container) and can use the `make` commands here (and only here... not outside the container)

- **Build the application in release mode**:
  ```bash
  make build
  ```
- **Build the application in debug mode**:
  ```bash
  make debug-build
  ```
- **Run the application**:
  ```bash
  make run
  ```
- **Run tests**:
  ```bash
  make test
  ```
- **Lint the application**:
  ```bash
  make lint
  ```
- **Format the code**:
  ```bash
  make format
  ```
- **Clean build artifacts**:
  ```bash
  make clean
  ```
- **Watch for changes and rebuild automatically**:
  ```bash
  make watch
  ```

---

## Notes
- The `Makefile` provides a streamlined workflow, ensuring that all operations are performed within the container to maintain consistency.
