# **Craft CLI Tool**


![Craft Logo](assets/logo.png)

## **Overview**

The `Craft` CLI is a tool designed to simplify the process of bootstrapping new projects by generating boilerplate code for various languages and frameworks. It ensures projects are configured to run seamlessly in Docker containers, enabling fast and consistent development environments. Whether you're working with Go, Java, or other supported languages, `Craft` helps you get started and maintain your projects efficiently.


---

## **Key Features**

- **Project Scaffolding** (`new` command):
  - Quickly generate project files and structure for supported languages and frameworks.
  - Embedded templates ensure a consistent starting point for new projects.

- **Project Updates** (`update` command):
  - Modify and maintain existing projects with additional boilerplate code or configuration updates.

- **Inspection** (`inspect` command):
  - View all allowed operations and their supported language/framework combinations.

- **Flexible Project Naming**:
  - Use the current directory name as the project name.
  - Specify a custom project name with a command-line flag.

- **Docker-Ready**:
  - Generated and updated projects are pre-configured to run in Docker containers.

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/VincentVanCode101/craft.git
   cd craft
   ```

2. Build the binary:
   ```bash
   ./getBinary.sh
   ```
  - the binary is automatically added to `/usr/local/bin/craft`
---

## **Usage**

### **1. Create a New Project**

Generate a new project by specifying the language:
```bash
craft new <language>
```

#### **Options**:
- `--name, -n`: Specify a name for the new project.
- `--current-dir-name, -c`: Use the current directory name for the project.

#### **Examples**:
1. Generate a Go project in the current directory:
   ```bash
   craft new go -c
   ```

2. Generate a new Java project with Maven in a directory named `MyJavaApp`:
   ```bash
   craft new java-maven-noframework -n MyJavaApp
   ```

3. View help for the `new` command:
   ```bash
   craft new --help
   ```

### **2. Update an Existing Project**

Modify an existing project to include additional boilerplate or updated configuration:
```bash
craft update <language>
```

#### **Examples**:
1. Update a Java project with a specific framework:
   ```bash
   craft update java-maven-quarkus
   ```

2. Update a Go project:
   ```bash
   craft update go
   ```

3. View help for the `update` command:
   ```bash
   craft update --help
   ```

---

### **3. Inspect Supported Operations**

View all allowed operations and their supported languages/framework combinations:
```bash
craft inspect
```

#### **Example**:
```bash
craft inspect
```
Example output:
```
Allowed Operations:
- new: go, java-maven-noframework
- update: go
```

---

## **How It Works**

1. The `new` command dynamically validates the language and framework.
2. Based on the selected language, the tool identifies and loads the appropriate embedded template.
3. The project files are generated in either:
   - A new directory (if `--name` is specified).
   - The current directory (if `--current-dir-name` is specified).
4. The generated project is ready to run in a Docker container.

---


## **Supported Languages**
Run the following command to view the list of supported languages:
```bash
craft inspect
```
### **Go**
- **Allowed Operations**:
  - `new`: Create a new Go project ([Documentation](./docs/go.md)).
  - `update`: Update an existing Go project ([Documentation](./docs/go.md)).

### **Java**
- **Allowed Build Tools and Frameworks**:
  - **Maven**:
    - `noframework`: Create a Java projects without any specific framework. ([Documentation](./docs/java-maven-noframework.md))
    - `spring`: Coming soon...
    - `quarkus`: Coming soon...
  - **Gradle**:
    - `noframework`: Coming soon...
    - `spring`: Coming soon...
    - `quarkus`: Coming soon...

## **Contributing**

1. Fork the repository.
2. Add support for new languages or frameworks by updating the `registry` and providing templates.
3. Submit a pull request with detailed information about the changes.

---

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

With `Craft`, you can spend less time setting up projects and more time building them. 🚀

