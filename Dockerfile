# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye as build-env

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates \
    # must install cross compiler for arm64
    gcc-aarch64-linux-gnu

WORKDIR /usr/src

# Get Go dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy rest of files
COPY . .

# compile dualityd to ARM64 architecture for final image
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    CGO_ENABLED=1 \
    CC=aarch64-linux-gnu-gcc \
    GOOS=linux \
    GOARCH=arm64 \
    go build -o build/dualityd_arm64 ./cmd/dualityd


# Final image build on small stable release of ARM64 Linux
FROM arm64v8/alpine:20220715 as base-env

# Install ca-certificates
RUN apk add --update \
    # required for dualityd to work
    libc6-compat \
    # allow JSON parsing in startup shell scripts
    jq \
    # required for HTTPS to connect properly
    ca-certificates

# Define network settings to be used (defined under top-level folder networks/)
ARG NETWORK=duality-1

# Make NETWORK available as an ENV variable for the running proccess
ENV NETWORK=$NETWORK

WORKDIR /usr/src

# Copy over binaries from the build-env
COPY --from=build-env /usr/src/build/dualityd_arm64 /usr/bin/dualityd

# expose ports
# rpc
EXPOSE 26657
# p2p
EXPOSE 26656
# grpc-web
EXPOSE 9091
# grpc
EXPOSE 9090
# prof
EXPOSE 6060
# api
EXPOSE 1317


FROM base-env as edit-config-files

# install TOML editing tool dasel for complicated TOML edits
RUN wget https://github.com/TomWright/dasel/releases/download/v1.27.3/dasel_linux_arm64.gz; \
    gzip -d dasel_linux_arm64.gz; \
    chmod 755 dasel_linux_arm64; \
    mv ./dasel_linux_arm64 /usr/local/bin/dasel;

#  create default config files
RUN dualityd init --chain-id "$NETWORK" duality

# edit config files
# determine some settings by either being a mainnet or testnet
ARG IS_MAINNET
# default testnet settings depending on network name (eg. duality-1, duality-13 are production chains)
RUN IS_MAINNET=${IS_MAINNET-$([[ "$NETWORK" =~ "^duality-\d+$" ]] && echo "true" || echo "")}; \
    # enable API to be served on any browser page when in developent, but only production web-app in production
    dasel put bool   -f /root/.duality/config/app.toml    ".api.enable" "true"; \
    dasel put bool   -f /root/.duality/config/app.toml    ".api.enabled-unsafe-cors" "$([[ $IS_MAINNET ]] && echo "false" || echo "true")"; \
    dasel put string -f /root/.duality/config/config.toml ".rpc.cors_allowed_origins" "$([[ $IS_MAINNET ]] && echo "app.duality.xyz" || echo "*")"; \
    # ensure listening to the RPC port doesn't block outgoing RPC connections
    dasel put string -f /root/.duality/config/config.toml ".rpc.laddr" "tcp://0.0.0.0:26657"; \
    # todo: add Prometheus telemetry
    # set chain id to network name
    dasel put string -f /root/.duality/config/client.toml ".chain-id" "$NETWORK";


# take configured files but don't take the dasel binary (the TOML files should be always safe to use)
FROM base-env as configured-env

# add configuration files for our binary
COPY --from=edit-config-files /root/.duality/config/app.toml /root/.duality/config/app.toml
COPY --from=edit-config-files /root/.duality/config/client.toml /root/.duality/config/client.toml
COPY --from=edit-config-files /root/.duality/config/config.toml /root/.duality/config/config.toml


# define continuation of configured-env as the default env
FROM configured-env

# add startup scripts and their dependencies
COPY scripts scripts
COPY networks networks

# default to serving the chain with default data and name
CMD ["sh", "./scripts/startup.sh"]
