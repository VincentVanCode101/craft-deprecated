# **Craft CLI Tool**


![Craft Logo](assets/logo.png)

## **Overview**
The `Craft` CLI is a tool designed to simplify the process of bootstrapping new projects by generating boilerplate code for various languages and frameworks.

- 🛠️ **Simplifies Project Bootstrapping**: Generates boilerplate code for various languages and frameworks.  
- 🐳 **Docker-Ready**: All projects are designed to run in Docker containers, allowing you to create and use them without installing the required language.
- 🌐 **Multi-Language Support**: Works with Go, Java, and other supported languages.  
- 🏎️ **Small and Fast**: A lightweight CLI tool designed for quick execution.  
- 💡 **Efficient and Reliable**: Helps you start and maintain projects effortlessly.  
- 🏗️ **almost no dependecies**: Requires only Docker to build and run your projects. There's no need to install any other language runtimes, compilers, or frameworks locally.

See why `Craft` solves the [problems of other scaffolding tools](#️-problems-with-other-scaffolding-tools) where I explain how it addresses common pitfalls and inefficiencies.

## **Table of Contents**

- [Problems with Other Scaffolding Tools](#️-problems-with-other-scaffolding-tools)
- [Why Use Craft](#️-why-use-craft)
- [Features](#-features)
- [Installation](#-installation)
- [Command Line Usage](#-command-line-usage)
  - [Creating New Projects](#1-creating-new-projects)
  - [Inspecting Supported Operations](#2-inspecting-supported-operations)
- [Supported Languages](#supported-languages)
- [License](#license)

---
## ⚠️ Problems with Other Scaffolding Tools

While exploring new languages or starting small projects, I encountered several issues with existing scaffolding tools:

1. **Language-Specific Tools**: Existing tools are often tied to a single language, limiting flexibility for multi-language workflows.  
2. **Language or Runtime-Specific Dependencies**: Most tools require the target language or its runtime to be installed on the machine, making setup cumbersome.
3. **Complex Projects**: These tools often generate large, interconnected setups that can be overwhelming for beginners or unnecessary for small tasks.  
4. **No Dockerized Setup**: Few tools create a containerized environment, making it harder to run the created projects in isolated and consistent environments.

These problems slow down productivity, create barriers for quickly experimenting with a new language or solving coding challenges, and result in projects that do not run in a containerized, isolated setup.
  - Running projects in Docker containers ensures consistency by providing a uniform environment across systems, eliminating 'it works on my machine' issues, isolating dependencies, and enabling easy cleanup or switching between projects without affecting the host system. ([see Docker](https://www.docker.com/))

---

## 🛠️ Why Use Craft?

- **Dockerized Development**: Automatically creates a containerized environment for every project, so you don’t need the language or runtime installed on your machine.
- **Lightweight and Fast**: Runs as a precompiled binary, ensuring fast execution without additional dependencies required on your host, expect for docker to setup some project dynamically.
- **Multi-Language Support**: Works seamlessly across multiple languages, making it versatile for various tasks.
- **Minimal Setup**: Generates only the essential files needed to start coding, with the option to create more complex setups if you’re familiar with the language. This allows you to build and structure your project the way you want.
- **Beginner-Friendly**: Focuses on simplicity and clarity, giving you exactly what you need to get started with a new language or task.

## ✨Features

- **Project Scaffolding** (`new` command):
  - Quickly generate project files and structure for supported languages and frameworks.
  - Automatically creates a new directory named `craft-<language>` for every project.
  - Specify dependencies for projects using the `--dependencies` flag.

- **Inspection** (`inspect` command):
  - View all allowed operations and their supported language/framework combinations.

- **Flexible Project Setup**:
  - Specify a custom project name using the `--name` flag.
  - View available dependencies for a specific language using the `--show-dependencies` flag.

- **Docker-Ready**:
  - Generated and updated projects are pre-configured to run in Docker containers.

---

## 📥 Installation

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

## 💻 Command Line Usage

### 1. Creating New Projects

Generate a new project by specifying the language:
```bash
craft new <language>
```

#### Options:
- `--name, -n`: Specify a name for the new project.
- `--dependencies, -d`: Specify the dependencies or project type (e.g., -d maven-spring).
- `--show-dependencies`: Show supported dependencies for the specified language.

#### Examples:
1. Generate a Go project named `craft-go`:
   ```bash
   craft new go
   ```

2. Generate a new Java project with Maven and Spring Boot in a directory named `MyJavaApp`:
   ```bash
   craft new java --name MyJavaApp --dependencies maven,spring
   ```
3. Show dependencies for the language `java`
- example output:
```
Maven is the default build-tool for the java projects:

Supported Dependencies:
  with the build tool 'maven' are:
    - Springboot
    - Quarkus
```

4. View help for the `new` command:
   ```bash
   craft new --help
   ```

---

### 2. Inspecting Supported Operations

View all allowed operations and their supported languages/framework combinations:
```bash
craft inspect
```
Example output:
```
Allowed Operations and Languages:
- Operation: New
  * Language: Java
    Run 'craft new java --help' to see the available dependencies.
  * Language: Go
    Run 'craft new go --help' to see the available dependencies.
```

---

# 🌐Supported Languages

<h2 style="display: flex; align-items: center; gap: 10px; font-size: 2rem; font-weight: bold; line-height: 1;">
  Go <img src="./assets/gopher.png" alt="Go Logo" style="height: 2rem;"/>
</h2>

<details>
<summary>more</summary>

#### Allowed Operations:
  - `new`: Create a new Go project ([Documentation](./docs/go.md)).
  
#### **Examples**:
1. Create a new Go project with the name `my-go-project`:
   ```bash
   craft new go --name my-go-project
   ```

2. Create a new Go project with the dependency `ncurses`:
   ```bash
   craft new go --name my-go-project --dependencies ncurses
   ```


</details>

<h2 style="display: flex; align-items: center; gap: 10px; font-size: 2rem; font-weight: bold; line-height: 1;">
  Java <img src="./assets/java.svg" alt="Go Logo" style="height: 2rem;"/>
</h2>

<details>
<summary>more</summary>

- **Allowed Build Tools and Frameworks**:
  - **Maven**:
    - `default`: Create a Java projects without any specific any dependencies will setup a plain java project with maven as the build tool. ([Documentation](./docs/java-maven-default.md))

    **Example**:
   Create a new Java project using Maven:
    ```bash
    craft new java -c
    ```
    - `springboot`: Coming soon...
    - `quarkus`: Coming soon...


</details>

---

## **📜License**

Licensed under [MIT License](./LICENSE)
