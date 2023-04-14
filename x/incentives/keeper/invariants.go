package keeper

// DONTCOVER

// // RegisterInvariants registers all governance invariants.
// func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
// 	ir.RegisterRoute(types.ModuleName, "accumulation-store-invariant", AccumulationStoreInvariant(keeper))
// 	ir.RegisterRoute(types.ModuleName, "stakes-amount-invariant", StakesBalancesInvariant(keeper))
// }

// // AccumulationStoreInvariant ensures that the sum of all stakeups at a given duration
// // is equal to the value stored within the accumulation store.
// func AccumulationStoreInvariant(keeper Keeper) sdk.Invariant {
// 	return func(ctx sdk.Context) (string, bool) {
// 		moduleAcc := keeper.ak.GetModuleAccount(ctx, types.ModuleName)
// 		balances := keeper.bk.GetAllBalances(ctx, moduleAcc.GetAddress())

// 		// check 1s, 1 day, 1 week, 2 weeks
// 		durations := []time.Duration{
// 			time.Second,
// 			time.Hour * 24,
// 			time.Hour * 24 * 7,
// 			time.Hour * 24 * 14,
// 		}

// 		// loop all denoms on stakeup module
// 		for _, coin := range balances {
// 			denom := coin.Denom
// 			for _, duration := range durations {
// 				accumulation := keeper.GetStakesAccumulation(ctx, types.QueryCondition{
// 					StakeQueryType: types.ByDuration,
// 					Denom:         denom,
// 					Duration:      duration,
// 				})

// 				stakes := keeper.GetStakesLongerThanDurationPair(ctx, denom, duration)
// 				stakeupSum := sdk.ZeroInt()
// 				for _, stake := range stakes {
// 					stakeupSum = stakeupSum.Add(stake.Coins.AmountOf(denom))
// 				}

// 				if !accumulation.Equal(stakeupSum) {
// 					return sdk.FormatInvariant(types.ModuleName, "accumulation-store-invariant",
// 						fmt.Sprintf("\taccumulation store value does not fit actual stakeup sum: %s != %s\n",
// 							accumulation.String(), stakeupSum.String(),
// 						)), true
// 				}
// 			}
// 		}

// 		return sdk.FormatInvariant(types.ModuleName, "accumulation-store-invariant", "All stakeup accumulation invariant passed"), false
// 	}
// }

// // StakesBalancesInvariant ensure that the module balance and the sum of all
// // tokens within all stakes have the equivalent amount of tokens.
// func StakesBalancesInvariant(keeper Keeper) sdk.Invariant {
// 	return func(ctx sdk.Context) (string, bool) {
// 		moduleAcc := keeper.ak.GetModuleAccount(ctx, types.ModuleName)
// 		balances := keeper.bk.GetAllBalances(ctx, moduleAcc.GetAddress())

// 		// loop all denoms on stakeup module
// 		for _, coin := range balances {
// 			denom := coin.Denom
// 			stakedAmount := sdk.ZeroInt()
// 			stakesByDenom := keeper.GetStakesPair(ctx, denom)
// 			for _, stake := range stakesByDenom {
// 				stakedAmount = stakedAmount.Add(stake.Coins.AmountOf(denom))
// 			}
// 			if !stakedAmount.Equal(coin.Amount) {
// 				return sdk.FormatInvariant(types.ModuleName, "stakes-amount-invariant",
// 					fmt.Sprintf("\tstakes amount of %s does not fit actual module balance: %s != %s\n",
// 						denom, stakedAmount.String(), coin.Amount.String(),
// 					)), true
// 			}
// 		}

// 		return sdk.FormatInvariant(types.ModuleName, "stakes-amount-invariant", "All stakeup amount invariant passed"), false
// 	}
// }
