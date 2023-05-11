#!/bin/sh

echo "Updating config files"

echo "Moniker: $NODE_MONIKER"

dasel put bool -f ${HOME}/.duality/config/app.toml -s ".api.enable" "true"
dasel put bool -f ${HOME}/.duality/config/app.toml -s ".api.enabled-unsafe-cors" "true"
dasel put string -f ${HOME}/.duality/config/config.toml -s ".rpc.cors_allowed_origins" "*"
dasel put bool   -f ${HOME}/.duality/config/config.toml -s ".p2p.addr_book_strict" "false"
# ensure listening to the RPC port doesn't block outgoing RPC connections
dasel put string -f ${HOME}/.duality/config/config.toml -s ".rpc.laddr" "tcp://0.0.0.0:26657"
dasel put string -f ${HOME}/.duality/config/client.toml -s ".chain-id" $CHAIN_ID

dasel put string -f ${HOME}/.duality/config/config.toml -s ".moniker" $NODE_MONIKER
