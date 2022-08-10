# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.238.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates \
    # gcc \
    # make \
    curl

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="v16.14.2"
RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install $NODE_VERSION" 2>&1

# [Choice] Ignite CLI version (align this with go.mod "github.com/ignite-hq/cli" version)
ARG IGNITE_CLI_VERSION="v0.22.0"
RUN curl "https://get.ignite.com/cli@$IGNITE_CLI_VERSION!" | bash

# see docs for exposed ports:
#   https://docs.ignite.com/kb/config.html#host
EXPOSE 26657
EXPOSE 26656
EXPOSE 6060
EXPOSE 9090
EXPOSE 1317
