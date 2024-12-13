# Craft Tool

`craft` is a CLI tool that generates boilerplate code for various programming languages using Docker. It is designed to be easy to set up and use, requiring minimal configuration.

## Quick Start

1. **Clone the Repository**
   ```bash
   git clone https://github.com/VincentVanCode101/craft.git
   cd craft
   ```

2. **Run the Setup Script**
   The setup script builds the Docker image and adds the `craft` script to your `PATH`. It also creates a symlink in `~/dotfiles/bin/.local/scripts` for convenience.
   ```bash
   ./setup.sh
   ```

3. **Use the Craft Tool**
   Generate boilerplate code for a new project with a single command. For example, to create a new Rust project:
   ```bash
   craft --operation=new --language=rust
   ```

   To display help information, simply call:
   ```bash
   craft
   ```

## How It Works

- The `setup.sh` script:
  - Builds the `craft` Docker image using the provided `Dockerfile`.
  - Creates a symbolic link to the `craft` script in `/usr/local/bin`, making it accessible from anywhere.
  - Creates an additional symbolic link in `~/dotfiles/bin/.local/scripts` for personal configuration.

- The `craft` command:
  - Runs the `craft` tool inside a lightweight Docker container.
  - Mounts your current directory into the container, ensuring that generated files are created on your host system.

## Features

- Supports multiple programming languages.
- Automatically generates boilerplate files like `Dockerfile`, `Makefile`, and others.
- Ensures generated files have the correct ownership and permissions on the host system.
