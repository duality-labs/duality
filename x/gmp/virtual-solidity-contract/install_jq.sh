#!/bin/bash

git clone https://github.com/stedolan/jq.git
cd jq
autoreconf -i
./configure --disable-maintainer-mode
make
sudo make install