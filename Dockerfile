FROM golang:1.23.3 AS dev

WORKDIR /app

COPY . .

FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o crafter .

FROM scratch

COPY --from=builder /app/crafter /usr/local/bin/crafter

USER 65532:65532

ENTRYPOINT ["crafter"]