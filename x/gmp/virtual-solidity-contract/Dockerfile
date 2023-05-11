# We'll use node:20 as the base image as truffle is a node.js application
FROM node:14-alpine

# Install jq
RUN apk update && apk add --no-cache curl jq python3 make g++

# Download and install the desired Solidity compiler version
RUN wget -q https://github.com/ethereum/solidity/releases/download/v0.8.4/solc-static-linux -O /usr/local/bin/solc \
    && chmod +x /usr/local/bin/solc

# These need to be separate
RUN npm install web3
RUN npm install --unsafe-perm -g truffle ganache-cli

# Create a directory for our application
WORKDIR /usr/src/app

# Copy our application into the docker image
COPY . .

RUN chmod +x /usr/src/app/startup.sh
RUN ls -la

# Run the startup script when the container starts
ENTRYPOINT ["/bin/sh", "-eu", "/usr/src/app/startup.sh"]