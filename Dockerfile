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
RUN go mod download

# Copy rest of files
COPY . .

# compile dualityd to ARM64 architecture for final image
RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o build/dualityd_arm64 ./cmd/dualityd


# Final image build on small stable release of ARM64 Linux
FROM arm64v8/alpine:20220715

# Install ca-certificates
RUN apk add --update \
    # required for dualityd to work
    libc6-compat \
    # required for HTTPS to connect properly
    ca-certificates

WORKDIR /usr/src

# Copy over binaries from the build-env
COPY --from=build-env /usr/src/build/dualityd_arm64 /usr/bin/dualityd

# create duality chain config files
RUN dualityd init duality

# add test accounts
RUN mkdir /root/.duality/testkeys
# alice
RUN dualityd keys add alice --keyring-backend test --output json > /root/.duality/testkeys/alice.json
RUN dualityd add-genesis-account $(dualityd keys show alice -a --keyring-backend test) 1000000000token,1000000000stake --keyring-backend test
# bob
RUN dualityd keys add bob --keyring-backend test --output json > /root/.duality/testkeys/bob.json
RUN dualityd add-genesis-account $(dualityd keys show bob -a --keyring-backend test) 1000000000token,1000000000stake --keyring-backend test

# Add gentxs to the genesis file
RUN dualityd gentx alice 1000000stake --chain-id duality --keyring-backend test
RUN dualityd collect-gentxs

# see docs for exposed ports:
#   https://docs.ignite.com/kb/config.html#host
EXPOSE 26657
EXPOSE 26656
EXPOSE 9090
EXPOSE 6060
EXPOSE 1317

# add startup scripts
COPY testnet/scripts scripts

# default to serving the chain with default data and name
CMD ["sh", "./scripts/startup.sh"]
