#!/bin/bash
set -e

# create new person with funds
# (token amounts here are measured in utoken denom)
person=$(bash /root/.duality/scripts/test_helpers.sh createAndFundUser 1000000000000stake,1000000000000token)

# add some helper functions to generate chain CLI args
count=100; # should be divisible by 4
function join_with_comma {
  local IFS=,
  echo "$*"
}
function repeat_with_comma {
  repeated=()
  for (( i=0; i<$count/2; i++ ))
  do
    repeated+=( $1 )
  done
  join_with_comma "${repeated[@]}"
}
function get_token_1_reserves_amount {
  amount=$1
  index=$2;
  echo " $amount / (1.0001 ^ $index) " \
    | bc -l \
    | awk '{printf("%d\n",$0+1)}' # round up only (in case we don't create enough reserves)
}

# create initial tick array outside of max price amplitude
max_tick_index=12000
indexes0=()
indexes1=()
amounts0=()
amounts1=()
amount=1000000
fee=30
for (( i=0; i<$count/4; i++ ))
do
  index=$(( $RANDOM % $max_tick_index ))
  indexes0+=( $(( -$max_tick_index - $index )) -$index )
  indexes1+=( $index $(( $index + $max_tick_index )) )
  # calculate reserve amounts to add that will equal the same amount of shares
  amounts0+=( $amount $amount )
  amounts1+=( $( get_token_1_reserves_amount $amount $index ) )
  amounts1+=( $( get_token_1_reserves_amount $amount $(( $index + $max_tick_index )) ) )
done
indexes=( "${indexes0[@]}" "${indexes1[@]}" )

# apply an amount to all tick indexes specified
dualityd tx dex deposit \
  $(dualityd keys show "$person" --output json | jq -r .address) \
  stake \
  token \
  "$(repeat_with_comma "$amount"),$(repeat_with_comma "0")" \
  "$(repeat_with_comma "0"),$(join_with_comma "${amounts1[@]}")" \
  "[$(join_with_comma "${indexes0[@]}"),$(join_with_comma "${indexes1[@]}")]" \
  "$(repeat_with_comma "$fee"),$(repeat_with_comma "$fee")" \
  "$(repeat_with_comma "false"),$(repeat_with_comma "false")" \
  --from $person --yes --output json --broadcast-mode block --gas 10000000 \
  | jq -r '"[ tx code: \(.code) ] [ tx hash \(.txhash) ]"' \
  | xargs -I{} echo "{} deposited: initial $count seed liquidity ticks"

# approximate price with sine curves of given amplitude and period
# macro curve oscillates over hours
amplitude1=10000 # in ticks
period1=3600 # in seconds
# micro curve oscillates over minutes
amplitude2=-2000 # in seconds
period2=300 # in seconds
two_pi=$( echo "scale=8; 8*a(1)" | bc -l )

# respond to price changes forever
while true
do
  # wait a bit, maybe less than a block or enough that we don't touch a block or two
  sleep $(( $RANDOM % 20 + 2 ))

  # determine the new current price goal
  current_price=$( \
    echo " $amplitude1*s($EPOCHSECONDS / $period1 * $two_pi) + $amplitude2*s($EPOCHSECONDS / $period2 * $two_pi) " \
    | bc -l \
    | awk '{printf("%d\n",$0+0.5)}' \
  )

  # add some randomness into price goal
  goal_price=$(( $current_price + $RANDOM % 1000 - 500 ))

  # - make a swap to get to current price

  # first, find the reserves of tokens that are outside the desired price
  # then swap those reserves
  reserves0=$( \
    wget -q -O - $API_ADDRESS/dualitylabs/duality/dex/tick_liquidity/stake%3C%3Etoken/stake?pagination.limit=100 \
    | jq "[.tickLiquidity[].poolReserves | select((.tickIndex | tonumber) > $goal_price) | .reserves | tonumber] | add as \$sum | if \$sum == null then 0 else \$sum end" \
  )
  if [[ $reserves0 -gt 0 ]]
  then
    dualityd tx dex swap \
      $(dualityd keys show "$person" --output json | jq -r .address) \
      $(( $reserves0 * 100 )) \
      token \
      stake \
      --max-amount-out $reserves0 \
      --from $person --yes --output json --broadcast-mode block --gas 10000000 \
      | jq -r '"[ tx code: \(.code) ] [ tx hash \(.txhash) ]"' \
      | xargs -I{} echo "{} swapped:   ticks toward target tick index of $goal_price"
  else
    reserves1=$( \
      wget -q -O - $API_ADDRESS/dualitylabs/duality/dex/tick_liquidity/stake%3C%3Etoken/token?pagination.limit=100 \
      | jq "[.tickLiquidity[].poolReserves | select((.tickIndex | tonumber) < $goal_price) | .reserves | tonumber] | add as \$sum | if \$sum == null then 0 else \$sum end" \
    )
    if [[ $reserves1 -gt 0 ]]
    then
      dualityd tx dex swap \
        $(dualityd keys show "$person" --output json | jq -r .address) \
        $(( $reserves1 * 100 )) \
        stake \
        token \
        --max-amount-out $reserves1 \
        --from $person --yes --output json --broadcast-mode block --gas 10000000 \
        | jq -r '"[ tx code: \(.code) ] [ tx hash \(.txhash) ]"' \
        | xargs -I{} echo "{} swapped:   ticks toward target tick index of $goal_price"
    fi
  fi

  # - replace the end pieces of liquidity with values closer to the current price

  # determine new indexes close to the current price
  new_index0=$(( $current_price - 1000 - $RANDOM % 1000 ))
  new_index1=$(( $current_price + 1000 + $RANDOM % 1000 ))

  # add these extra ticks to prevent swapping though all ticks errors
  # we deposit first to lessen the cases where we have entirely one-sided liquidity
  dualityd tx dex deposit \
    $(dualityd keys show "$person" --output json | jq -r .address) \
    stake \
    token \
    "$amount,0" \
    "0,$( get_token_1_reserves_amount $amount $new_index1 )" \
    "[$new_index0,$new_index1]" \
    "$fee,$fee" \
    false,false \
    --from $person --yes --output json --broadcast-mode block --gas 300000 \
    | jq -r '"[ tx code: \(.code) ] [ tx hash \(.txhash) ]"' \
    | xargs -I{} echo "{} deposited: new close-to-price ticks $new_index0, $new_index1"

  # add to array
  indexes+=( $new_index0 $new_index1 )
  # then sort array
  indexes=( $( printf '%s\n' "${indexes[@]}" | sort -n ) )

  # withdraw the end values
  dualityd tx dex withdrawal \
    $(dualityd keys show "$person" --output json | jq -r .address) \
    stake \
    token \
    "$amount,$amount" \
    "[$(( indexes[0] )),$(( indexes[-1] ))]" \
    "$fee,$fee" \
    --from $person --yes --output json --broadcast-mode block --gas 300000 \
    | jq -r '"[ tx code: \(.code) ] [ tx hash \(.txhash) ]"' \
    | xargs -I{} echo "{} withdrew:  end ticks $(( indexes[0] )), $(( indexes[-1] ))"

  # remove the withdrawn indexes
  unset 'indexes[0]'
  unset 'indexes[-1]'

done
