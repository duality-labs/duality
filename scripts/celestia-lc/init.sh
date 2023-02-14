#!/bin/sh

/celestia light init \
    --p2p.network $P2P_NETWORK --node.store $NODE_STORE
(echo $KEY_PASSWD; echo $KEY_PASSWD) | /cel-key add $KEY_NAME \
    --keyring-backend $KEYRING_BACKEND --keyring-dir $NODE_STORE/keys \
    --node.type light --p2p.network $P2P_NETWORK
