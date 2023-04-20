[View code on GitHub](https://github.com/duality-labs/duality/incentives/types/gauge.go)

The `types` package contains the `Gauge` struct and associated methods used in the duality project. The `Gauge` struct represents a gauge that tracks the distribution of rewards over a set number of epochs. The `NewGauge` function creates a new `Gauge` struct with the specified parameters. The `hasEpochsRemaining` method returns a boolean indicating whether the gauge has any epochs remaining to distribute rewards. The `hasStarted` method returns a boolean indicating whether the gauge has started distributing rewards. The `IsUpcomingGauge` method returns a boolean indicating whether the gauge is upcoming, i.e., its distribution start time is after the provided time. The `IsActiveGauge` method returns a boolean indicating whether the gauge is currently active, i.e., it has started distributing rewards and has epochs remaining. The `IsFinishedGauge` method returns a boolean indicating whether the gauge has finished distributing rewards. The `RewardsNextEpoch` method returns the rewards that will be distributed in the next epoch. The `EpochsRemaining` method returns the number of epochs remaining for the gauge to distribute rewards. The `CoinsRemaining` method returns the coins remaining to be distributed by the gauge.

This code is used to manage the distribution of rewards over a set number of epochs. The `Gauge` struct represents a gauge that tracks the distribution of rewards. The `NewGauge` function is used to create a new gauge with the specified parameters. The `hasEpochsRemaining`, `hasStarted`, `IsUpcomingGauge`, `IsActiveGauge`, and `IsFinishedGauge` methods are used to determine the state of the gauge, i.e., whether it is upcoming, active, or finished. The `RewardsNextEpoch` method is used to determine the rewards that will be distributed in the next epoch. The `EpochsRemaining` method is used to determine the number of epochs remaining for the gauge to distribute rewards. The `CoinsRemaining` method is used to determine the coins remaining to be distributed by the gauge.

Example usage:

```
gauge := NewGauge(1, true, QueryCondition{}, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, time.Now(), 10, 5, sdk.Coins{}, 1)
if gauge.IsUpcomingGauge(time.Now()) {
    fmt.Println("Gauge is upcoming")
} else if gauge.IsActiveGauge(time.Now()) {
    fmt.Println("Gauge is active")
} else if gauge.IsFinishedGauge(time.Now()) {
    fmt.Println("Gauge is finished")
}
rewards := gauge.RewardsNextEpoch()
fmt.Println("Rewards next epoch:", rewards.String())
epochsRemaining := gauge.EpochsRemaining()
fmt.Println("Epochs remaining:", epochsRemaining)
coinsRemaining := gauge.CoinsRemaining()
fmt.Println("Coins remaining:", coinsRemaining.String())
```
## Questions: 
 1. What is the purpose of the `Gauge` struct and what are its required parameters?
- The `Gauge` struct is used to represent a gauge and its parameters include an ID, whether it is perpetual, a query condition for distribution, coins, start time, number of epochs paid over, filled epochs, distributed coins, and pricing tick.

2. What is the difference between `IsUpcomingGauge`, `IsActiveGauge`, and `IsFinishedGauge`?
- `IsUpcomingGauge` returns true if the gauge's distribution start time is after the provided time, `IsActiveGauge` returns true if the gauge is in an active state during the provided time, and `IsFinishedGauge` returns true if the gauge is in a finished state during the provided time.

3. What is the purpose of the `RewardsNextEpoch`, `EpochsRemaining`, and `CoinsRemaining` functions?
- `RewardsNextEpoch` calculates the rewards for the next epoch based on the remaining coins and epochs, `EpochsRemaining` calculates the number of epochs remaining based on whether the gauge is perpetual or not, and `CoinsRemaining` calculates the remaining coins based on the coins and distributed coins.