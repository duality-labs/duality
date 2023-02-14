#!/bin/sh

/celestia light start\
	--p2p.network $P2P_NETWORK \
	--core.ip $NODE_ADDR --gateway --gateway.addr $GATEWAY_ADDR --gateway.port $GATEWAY_PORT
