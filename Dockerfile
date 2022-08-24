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
EXPOSE 1317

# default to serving the chain with default data and name
CMD ["dualityd", "start", "--moniker", "duality"]
