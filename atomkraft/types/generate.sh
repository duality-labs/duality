#!/bin/bash -ex

BASE_DIR=$(dirname $(readlink -f "$0"))
PROJECT_DIR=$(dirname $(dirname $BASE_DIR))
PROTO_DIR=$PROJECT_DIR/proto

npm install &> /dev/null
for f in $(find $PROTO_DIR -name "*.proto"); do
  protoc \
    -I=$PROTO_DIR \
    -I=$BASE_DIR/node_modules/@protobufs \
    --python_out=$BASE_DIR \
    $f
done