#!/bin/bash
set -e

createAndFundUser() {
    tokens=$1
    # create person name
    person=$(openssl rand -hex 12)
    # create person's new account
    dualityd keys add $person <<< $'asdfasdf\nn' >/dev/null
    # use Fred Faucet's funds to fund their account
    dualityd tx bank send $(dualityd keys show fred --output json | jq -r .address) $(dualityd keys show $person --output json | jq -r .address) $tokens -y --broadcast-mode block >/dev/null

    echo "$person"
}


# below code is taken from https://stackoverflow.com/questions/8818119/how-can-i-run-a-function-from-a-script-in-command-line#16159057

# Check if the function exists (bash specific)
if declare -f "$1" > /dev/null
then
  # call arguments verbatim
  "$@"
else
  # Show a helpful error
  echo "'$1' is not a known function name" >&2
  exit 1
fi
