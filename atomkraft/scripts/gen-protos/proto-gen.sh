#!/bin/bash -ex

pip install betterproto[compiler]
npm install

SCRIPT_DIR=$(dirname $(readlink -f "$0"))
AK_CONFIG_DIR=$(dirname $(dirname $SCRIPT_DIR))
REPO_DIR=$(dirname $AK_CONFIG_DIR)
PROTO_DIR=$REPO_DIR/proto
DEP_PROTO_DIR=$SCRIPT_DIR/node_modules/@protobufs
OUT_DIR=$AK_CONFIG_DIR/reactors/proto

npm install &> /dev/null
protoc \
  -I=$PROTO_DIR \
  -I=$SCRIPT_DIR/node_modules/@protobufs \
  --python_betterproto_out=$OUT_DIR \
  $(find $PROTO_DIR $DEP_PROTO_DIR -path -prune -o -name "*.proto" -print0 | xargs -0)