FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o craft .

FROM debian:bullseye-slim AS runtime

WORKDIR /app

COPY --from=builder /app/craft /usr/local/bin/craft

VOLUME /app

ENTRYPOINT ["craft"]
