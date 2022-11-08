# NOTE: Since we are copying over the core project deps we need to run docker build from the root directory
# ie. docker build -f ./testnet-faucet/Dockerfile -t faucet .

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye as build-env

ARG PORT=9000
ENV PORT=$PORT

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
#    build-essential \
#    ca-certificates \
    # must install cross compiler for arm64
    gcc-aarch64-linux-gnu

WORKDIR /usr/src

# Get Go dependencies
# TODO: this is grabbing all deps from core project, we don't really need this
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

# Copy work files for faucet (but not Dockerfile: to avoid invalidating caches while developing)
COPY testnet-faucet/faucet ./faucet

# Copy work files for dualityd (but not Dockerfile: to avoid invalidating caches while developing)
COPY app ./app
COPY cmd ./cmd
COPY docs ./docs
COPY proto ./proto
COPY testutil ./testutil
COPY vue ./vue
COPY x ./x

# compile faucet to ARM64 architecture for final image
RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o build/faucet_arm64 ./faucet/faucet.go

RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o build/dualityd_arm64 ./cmd/dualityd


# Final image build on small stable release of ARM64 Linux
FROM arm64v8/alpine:20220715




 RUN apk add --update \
    # required for dualityd to work
    libc6-compat 
    # allow JSON parsing in startup shell scripts
    # jq \
    # required for HTTPS to connect properly
    # ca-certificates

WORKDIR /usr/src

# Copy over binaries from the build-env
COPY --from=build-env /usr/src/build/faucet_arm64 /usr/bin/faucet
COPY --from=build-env /usr/src/build/dualityd_arm64 /usr/bin/dualityd

EXPOSE $PORT

COPY testnet-faucet/scripts scripts

# Start faucet 
CMD ["sh", "./scripts/startup.sh"]
