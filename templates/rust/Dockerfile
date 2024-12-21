# Use the official Rust image
FROM rust:latest

# Set the working directory in the container
WORKDIR /app

# Install make
RUN apt-get update && apt-get install -y make && rm -rf /var/lib/apt/lists/*

# Copy only the Cargo.toml and Cargo.lock to cache dependencies
COPY Cargo.toml Cargo.lock ./

# Pre-download dependencies
RUN cargo fetch

# Copy the source code into the container
COPY . .

# Set the default command to run when the container starts
CMD ["cargo", "run"]
