#!/bin/sh

# reset existing chain
dualityd tendermint unsafe-reset-all
# init chain
dualityd init $MONIKER --chain-id $CHAIN_ID

dasel put bool   -f /root/.duality/config/app.toml    ".api.enable" "true"; \
dasel put bool   -f /root/.duality/config/app.toml    ".api.enabled-unsafe-cors" "true"; \
dasel put string -f /root/.duality/config/config.toml ".rpc.cors_allowed_origins" "*"; \
    # ensure listening to the RPC port doesn't block outgoing RPC connections
dasel put string -f /root/.duality/config/config.toml ".rpc.laddr" "tcp://0.0.0.0:26657"; \

