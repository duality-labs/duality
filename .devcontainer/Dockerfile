# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.238.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.19-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    curl

# [Choice] Ignite CLI version (align this with go.mod "github.com/ignite-hq/cli" version)
ARG IGNITE_CLI_VERSION="v0.23.0"
RUN curl "https://get.ignite.com/cli@$IGNITE_CLI_VERSION!" | bash

RUN go install github.com/segmentio/golines@latest

# see docs for exposed ports:
#   https://docs.ignite.com/kb/config.html#host
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
