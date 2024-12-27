FROM rust:latest AS dev

WORKDIR /workspace

RUN apt-get update && apt-get install -y make && rm -rf /var/lib/apt/lists/*

RUN rustup component add clippy rustfmt \
    && cargo install cargo-watch

COPY Cargo.toml ./
COPY src ./src

RUN cargo fetch

COPY Makefile .

CMD ["cargo", "run"]
