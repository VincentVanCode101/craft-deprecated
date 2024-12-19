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
  Look for a container named `PROJECT_NAME-java-env`.

#### **2. Connect to the Development Container**
Once the container is running, connect to it for development purposes.

- **Open a bash session in the container:**
  ```bash
  docker exec -it PROJECT_NAME-java-env bash
  ```
  - use the `make` command from inside the container (see the chapter below)


### **How to Use the Makefile (Container Usage)**

This `Makefile` is designed to streamline the process of building, running, testing, and cleaning up a Java project inside a Docker container environment. The commands are optimized to work with a typical Java/Maven project structure and can be executed within the container.

---

### **Commands Overview**

#### **1. Default Target: `make` or `make all`**
- **Purpose**: Compiles all `.java` files in the `$(SOURCE_DIR)` (`src/main/java`) directory.
- **Usage**:
  ```bash
  make
  ```
- **Effect**:
  - Creates the `build` directory (if it doesn’t already exist).
  - Compiles all `.java` files into `$(BUILD_DIR)`.

---

#### **2. Build: `make build`**
- **Purpose**: Builds the project, either the entire project or a specific Java file.
- **Usage**:
  - **Build the entire project**:
    ```bash
    make build
    ```
  - **Build a specific file**:
    ```bash
    make build ARGS=src/main/java/com/main/foo/bar.java
    ```
- **Effect**:
  - Compiles all `.java` files into `$(BUILD_DIR)` when `ARGS` is not specified.
  - If `ARGS` is provided, only the specified file is compiled into `$(BUILD_DIR)`.

---

#### **3. Run: `make run`**
- **Purpose**: Runs the application, either the main class or a standalone file.
- **Usage**:
  - **Run the main project**:
    ```bash
    make run
    ```
  - **Run a specific file**:
    ```bash
    make run ARGS=src/main/java/com/main/foo/bar.java
    ```
- **Effect**:
  - Runs the `MAIN_CLASS` (defined as `com.main.App`) if `ARGS` is not specified.
  - If `ARGS` is provided, it derives the fully qualified class name from the file path and executes it.

---

#### **4. Create an Uber JAR: `make uber-jar`**
- **Purpose**: Creates an executable JAR file containing all compiled classes and the main class defined in the `MAIN_CLASS` variable.
- **Usage**:
  ```bash
  make uber-jar
  ```
- **Effect**:
  - Compiles all `.java` files (if not already compiled).
  - Generates an uber JAR named `PROJECT_NAME.jar` in the `$(JAR_DIR)` directory.
  - The JAR includes:
    - All compiled classes.
    - A manifest file with the `Main-Class` specified as `MAIN_CLASS`.

- **Example**:
  ```bash
  make uber-jar
  ```
  Output:
  ```
  Compiling all files in src/main/java...
  Creating uber JAR for PROJECT_NAME...
  Uber JAR created: jar/PROJECT_NAME.jar
  ```

---

#### **5. Run the Uber JAR: `make run-jar`**
- **Purpose**: Runs the uber JAR created by `make uber-jar`.
- **Usage**:
  ```bash
  make run-jar
  ```
- **Effect**:
  - Executes the `PROJECT_NAME.jar` file in the `$(JAR_DIR)` directory.
  - Automatically builds the uber JAR if it doesn’t already exist.

- **Example**:
  ```bash
  make run-jar
  ```
  Output:
  ```
  Running uber JAR for PROJECT_NAME without arguments...
  Hello World!
  ```

- **Pass Arguments to the JAR**:
  - Modify the `run-jar` command to pass arguments using the `ARGS` variable:
    ```bash
    make run-jar ARGS="arg1 arg2"
    ```
  - Example:
    ```bash
    make run-jar ARGS="foo bar"
    ```
  Output:
  ```
  Running uber JAR for PROJECT_NAME with arguments: foo bar...
  Hello World!
  ```
> [!NOTE]
> The arguments of "foo" and "bar" do not appear in the programm output since the App.java does not evaluate them... but if you had logic depending on the args, this run command would pass them correctly to the PROJECT_NAME.jar

---

#### **6. Run All Tests: `make test`**
- **Purpose**: Executes all tests using Maven.
- **Usage**:
  ```bash
  make test
  ```
- **Effect**:
  - Runs `mvn test`, executing all test cases defined in the project.

---

#### **7. Clean: `make clean`**
- **Purpose**: Cleans up build artifacts.
- **Usage**:
  ```bash
  make clean
  ```
- **Effect**:
  - Deletes the `$(BUILD_DIR)` and `$(JAR_DIR)` directories.

---

### **Best Practices**
- **Main Class**: Update the `MAIN_CLASS` variable in the `Makefile` if your main application class is different from `com.main.App`.
- **Project Name**: Update the `JAR_NAME` variable to reflect your application name.

This `Makefile` simplifies project management inside the container, enabling you to compile, run, package, and clean up your Java project efficiently.

---
