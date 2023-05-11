#!/bin/bash -eu

docker build -t truffle-ganache . &> /dev/null
docker run -it --rm truffle-ganache