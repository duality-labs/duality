# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates \
    # gcc \
    # make \
    curl

# [Choice] Ignite CLI version (align this with go.mod "github.com/ignite-hq/cli" version)
ARG IGNITE_CLI_VERSION="v0.22.0"
RUN curl "https://get.ignite.com/cli@$IGNITE_CLI_VERSION!" | bash

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="16.14.2"
RUN curl -fsSL https://deb.nodesource.com/setup_$NODE_VERSION | bash -
RUN apt-get install -y nodejs

WORKDIR /usr/src

# Get Go dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

# Copy rest of files
COPY . .

# compile dualityd
RUN go install ./cmd/dualityd

# create genesis files
RUN dualityd init duality

# see docs for exposed ports:
#   https://docs.ignite.com/kb/config.html#host
EXPOSE 26657
EXPOSE 26656
EXPOSE 9090
EXPOSE 6060
EXPOSE 4500
EXPOSE 1317

# default to serving the chain with default data and name
CMD ["ignite", "chain", "serve", "-r"]
