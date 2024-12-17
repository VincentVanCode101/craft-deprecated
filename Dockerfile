FROM golang:1.23.3 AS dev
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install golang.org/x/lint/golint@latest

COPY . .

RUN make linux-build

FROM busybox:1.37.0 AS runtime
WORKDIR /

COPY --from=dev /app/craft-native /usr/local/bin/craft-native

USER 65532:65532

ENTRYPOINT ["craft-native"]