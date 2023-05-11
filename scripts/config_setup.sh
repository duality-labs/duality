#!/bin/sh

echo "Updating config files"
# determine some settings by either being a mainnet or testnet


# enable API to be served on any browser page when in developent, but only production web-app in production
dasel put bool -f ${HOME}/.duality/config/app.toml -s ".api.enable" "true"
dasel put bool -f ${HOME}/.duality/config/app.toml -s ".api.enabled-unsafe-cors" "$([[ $IS_MAINNET = true ]] && echo "false" || echo "true")"
dasel put string -f ${HOME}/.duality/config/config.toml -s ".rpc.cors_allowed_origins" "$([[ $IS_MAINNET = true ]] && echo "app.duality.xyz" || echo "*")"

# if not mainnet this may be a localnet, where we need address book to not be strict
dasel put bool   -f ${HOME}/.duality/config/config.toml -s ".p2p.addr_book_strict" "false"
# ensure listening to the RPC port doesn't block outgoing RPC connections
dasel put string -f ${HOME}/.duality/config/config.toml -s ".rpc.laddr" "tcp://0.0.0.0:26657"
dasel put string -f ${HOME}/.duality/config/client.toml -s ".chain-id" $CHAIN_ID


