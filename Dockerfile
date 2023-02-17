FROM golang:1.18 as build-env

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates

WORKDIR /usr/src

# Get Go dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy rest of files
COPY . .

# build duality
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    CGO_ENABLED=1 \
    GOOS=linux \
    go build -o build/dualityd ./cmd/dualityd


FROM alpine:3.17 as base-env

RUN apk add --update \
    # required for dualityd to work
    libc6-compat \
    # allow JSON parsing in startup shell scripts
    jq \
    # required for HTTPS to connect properly
    ca-certificates \
    curl

# Copy over binaries and genesis files from the build-env
RUN wget https://github.com/TomWright/dasel/releases/download/v1.27.3/dasel_linux_arm64.gz; \
    gzip -d dasel_linux_arm64.gz; \
    chmod 755 dasel_linux_arm64; \
    mv ./dasel_linux_arm64 /usr/local/bin/dasel;
COPY --from=build-env /usr/src/build/dualityd /usr/bin/dualityd
COPY networks networks
COPY scripts/duality scripts
