#!/bin/sh -e

# place limit order
find . -type f -name "*.go" | xargs sed -i '' 's/PlacesLimitOrder/LimitBuys/g'
find . -type f -name "*.go" | xargs sed -i '' 's/PlacesLimitBuyOrder/LimitBuys/g'
find . -type f -name "*.go" | xargs sed -i '' 's/placesLimitOrder/limitBuys/g'

# cancel limit order
find . -type f -name "*.go" | xargs sed -i '' 's/CancelsLimitOrder/CancelsLimitSell/g'
find . -type f -name "*.go" | xargs sed -i '' 's/cancelsLimitOrder/cancelsLimitSell/g'

# swap
find . -type f -name "*.go" | xargs sed -i '' 's/PlacesSwapOrder/Buys/g'
find . -type f -name "*.go" | xargs sed -i '' 's/placesSwapOrder/buys/g'

# withdraw limit order
find . -type f -name "*.go" | xargs sed -i '' 's/WithdrawsFilledLimitOrder/WithdrawsLimitBuy/g'
find . -type f -name "*.go" | xargs sed -i '' 's/withdrawsFilledLimitOrder/withdrawsLimitBuy/g'
