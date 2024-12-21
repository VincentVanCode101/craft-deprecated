FROM rust:alpine AS builder

ARG UID=1000
ARG GID=1000
ARG PROJECT_NAME=default-project-name

WORKDIR /build-space

RUN cargo new ${PROJECT_NAME}

RUN chown -R 1000:1000 /build-space