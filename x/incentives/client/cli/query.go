package cli

import (
	"github.com/spf13/cobra"

	"github.com/duality-labs/duality/osmoutils/osmocli"
	"github.com/duality-labs/duality/x/incentives/types"
)

// GetQueryCmd returns the query commands for this module.
func GetQueryCmd() *cobra.Command {
	// group incentives queries under a subcommand
	cmd := osmocli.QueryIndexCmd(types.ModuleName)
	// qcGetter := types.NewQueryClient
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGauges)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdToDistributeCoins)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdGetGaugeByID)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdActiveGauges)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdActiveGaugesPerDenom)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdUpcomingGauges)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdUpcomingGaugesPerDenom)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdModuleBalance)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdModuleStakedAmount)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountUnstakingCoins)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountStakedPastTime)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountStakedPastTimeNotUnstakingOnly)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdTotalStakedByDenom)
	cmd.AddCommand(
	// GetCmdRewardsEstimate(),
	// GetCmdAccountUnstakeableCoins(),
	// GetCmdAccountStakedCoins(),
	// GetCmdAccountUnstakedBeforeTime(),
	// GetCmdAccountStakedPastTimeDenom(),
	// GetCmdStakedByID(),
	// GetCmdAccountStakedLongerDuration(),
	// GetCmdAccountStakedLongerDurationNotUnstakingOnly(),
	// GetCmdAccountStakedLongerDurationDenom(),
	// GetCmdOutputStakesJson(),
	// GetCmdAccountStakedDuration(),
	// GetCmdNextStakeID(),
	// osmocli.GetParams[*types.QueryParamsRequest](
	// 	types.ModuleName, types.NewQueryClient),
	)

	return cmd
}

// // GetCmdGauges returns all available gauges.
// func GetCmdGauges() (*osmocli.QueryDescriptor, *types.GetGaugesActiveUpcomingRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "gauges",
// 		Short: "Query all available gauges",
// 		Long:  "{{.Short}}",
// 	}, &types.GetGaugesActiveUpcomingRequest{}
// }

// // GetCmdToDistributeCoins returns coins that are going to be distributed.
// func GetCmdToDistributeCoins() (*osmocli.QueryDescriptor, *types.GetModuleCoinsToBeDistributedRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "to-distribute-coins",
// 		Short: "Query coins that is going to be distributed",
// 		Long:  `{{.Short}}`}, &types.GetModuleCoinsToBeDistributedRequest{}
// }

// // GetCmdGetGaugeByID returns a gauge by ID.
// func GetCmdGetGaugeByID() (*osmocli.QueryDescriptor, *types.GetGaugeByIDRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "gauge-by-id [id]",
// 		Short: "Query gauge by id.",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} gauge-by-id 1
// `}, &types.GetGaugeByIDRequest{}
// }

// // GetCmdActiveGauges returns active gauges.
// func GetCmdActiveGauges() (*osmocli.QueryDescriptor, *types.ActiveGetGaugesActiveUpcomingRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "active-gauges",
// 		Short: "Query active gauges",
// 		Long:  `{{.Short}}`}, &types.ActiveGetGaugesActiveUpcomingRequest{}
// }

// // GetCmdActiveGaugesPerDenom returns active gauges for a specified denom.
// func GetCmdActiveGaugesPerDenom() (*osmocli.QueryDescriptor, *types.ActiveGaugesPerDenomRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "active-gauges-per-den [den]denom [denom]",
// 		Short: "Query active gauges per denom",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} active-gauges-per-denom gamm/pool/1`}, &types.ActiveGaugesPerDenomRequest{}
// }

// // GetCmdUpcomingGauges returns scheduled gauges.
// func GetCmdUpcomingGauges() (*osmocli.QueryDescriptor, *types.UpcomingGetGaugesActiveUpcomingRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "upcoming-gauges",
// 		Short: "Query upcoming gauges",
// 		Long:  `{{.Short}}`}, &types.UpcomingGetGaugesActiveUpcomingRequest{}
// }

// // GetCmdUpcomingGaugesPerDenom returns scheduled gauges for specified denom..
// func GetCmdUpcomingGaugesPerDenom() (*osmocli.QueryDescriptor, *types.UpcomingGaugesPerDenomRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "upcoming-gauges-per-denom [denom]",
// 		Short: "Query scheduled gauges per denom",
// 		Long:  `{{.Short}}`}, &types.UpcomingGaugesPerDenomRequest{}
// }

// // GetCmdRewardsEstimate returns rewards estimation.
// func GetCmdRewardsEstimate() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "rewards-estimation",
// 		Short: "Query rewards estimation",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query rewards estimation.

// Example:
// $ %s query incentives rewards-estimation
// `,
// 				version.AppName,
// 			),
// 		),
// 		Args: cobra.ExactArgs(0),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			var res *types.AccountStakedLongerDurationResponse
// 			ownerStakes := []uint64{}
// 			stakeIds := []uint64{}

// 			owner, err := cmd.Flags().GetString(FlagOwner)
// 			if err != nil {
// 				return err
// 			}

// 			stakeIdsCombined, err := cmd.Flags().GetString(FlagStakeIds)
// 			if err != nil {
// 				return err
// 			}
// 			stakeIdStrs := strings.Split(stakeIdsCombined, ",")

// 			endEpoch, err := cmd.Flags().GetInt64(FlagEndEpoch)
// 			if err != nil {
// 				return err
// 			}

// 			// if user doesn't provide at least one of the stake ids or owner, we don't have enough information to proceed.
// 			if stakeIdsCombined == "" && owner == "" {
// 				return fmt.Errorf("either one of owner flag or stake IDs must be provided")

// 				// if user provides stakeIDs, use these stakeIDs in our rewards estimation
// 			} else if owner == "" {
// 				for _, stakeIdStr := range stakeIdStrs {
// 					stakeId, err := strconv.ParseUint(stakeIdStr, 10, 64)
// 					if err != nil {
// 						return err
// 					}
// 					stakeIds = append(stakeIds, stakeId)
// 				}

// 				// if no stakeIDs are provided but an owner is provided, we query the rewards for all of the stakes the owner has
// 			} else if stakeIdsCombined == "" {
// 				stakeIds = append(stakeIds, ownerStakes...)
// 			}

// 			// if stakeIDs are provided and an owner is provided, only query the stakeIDs that are provided
// 			// if a stakeID was provided and it doesn't belong to the owner, return an error
// 			if len(stakeIds) != 0 && owner != "" {
// 				for _, stakeId := range stakeIds {
// 					validInputStakeId := contains(ownerStakes, stakeId)
// 					if !validInputStakeId {
// 						return fmt.Errorf("stake-id %v does not belong to %v", stakeId, owner)
// 					}
// 				}
// 			}

// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)

// 			if owner != "" {
// 				queryClientStakeup := types.NewQueryClient(clientCtx)

// 				res, err = queryClientStakeup.AccountStakedLongerDuration(cmd.Context(), &types.AccountStakedLongerDurationRequest{Owner: owner, Duration: time.Millisecond})
// 				if err != nil {
// 					return err
// 				}
// 				for _, stakeId := range res.Stakes {
// 					ownerStakes = append(ownerStakes, stakeId.ID)
// 				}
// 			}

// 			// TODO: Fix accumulation store bug. For now, we return a graceful error when attempting to query bugged gauges
// 			RewardsEstimateResult, err := queryClient.RewardsEstimate(cmd.Context(), &types.RewardsEstimateRequest{
// 				Owner:    owner, // owner is used only when stakeIds are empty
// 				StakeIds:  stakeIds,
// 				EndEpoch: endEpoch,
// 			})
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(RewardsEstimateResult)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)
// 	cmd.Flags().String(FlagOwner, "", "Owner to receive rewards, optionally used when stake-ids flag is NOT set")
// 	cmd.Flags().String(FlagStakeIds, "", "the stake ids to receive rewards, when it is empty, all stake ids of the owner are used")
// 	cmd.Flags().Int64(FlagEndEpoch, 0, "the end epoch number to participate in rewards calculation")

// 	return cmd
// }

func contains(s []uint64, value uint64) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

// // GetCmdModuleBalance returns full balance of the stakeup module.
// // Stakeup module is where coins of stakes are held.
// // This includes staked balance and unstaked balance of the module.
// func GetCmdModuleBalance() (*osmocli.QueryDescriptor, *types.ModuleBalanceRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "module-balance",
// 		Short: "Query module balance",
// 		Long:  `{{.Short}}`}, &types.ModuleBalanceRequest{}
// }

// // GetCmdModuleStakedAmount returns staked balance of the module,
// // which are all the tokens not unstaking + tokens that are not finished unstaking.
// func GetCmdModuleStakedAmount() (*osmocli.QueryDescriptor, *types.ModuleStakedAmountRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "module-staked-amount",
// 		Short: "Query staked amount",
// 		Long:  `{{.Short}}`}, &types.ModuleStakedAmountRequest{}
// }

// // GetCmdAccountUnstakeableCoins returns unstakeable coins which has finsihed unstaking.
// // TODO: DELETE THIS + Actual query in subsequent PR
// func GetCmdAccountUnstakeableCoins() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "account-unstakeable-coins <address>",
// 		Short: "Query account's unstakeable coins",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query account's unstakeable coins.

// Example:
// $ %s query stakeup account-unstakeable-coins <address>
// `,
// 				version.AppName,
// 			),
// 		),
// 		Args: cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)

// 			res, err := queryClient.AccountUnstakeableCoins(cmd.Context(), &types.AccountUnstakeableCoinsRequest{Owner: args[0]})
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(res)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }

// // GetCmdAccountUnstakingCoins returns unstaking coins of a specific account.
// func GetCmdAccountUnstakingCoins() (*osmocli.QueryDescriptor, *types.AccountUnstakingCoinsRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-unstaking-coins <address>",
// 		Short: "Query account's unstaking coins",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-unstaking-coins <address>`}, &types.AccountUnstakingCoinsRequest{}
// }

// // GetCmdAccountStakedCoins returns staked coins that that are still in a staked state from the specified account.
// func GetCmdAccountStakedCoins() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedCoinsRequest](
// 		"account-staked-coins <address>",
// 		"Query account's staked coins",
// 		`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-coins <address>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedPastTime returns stakes of an account with unstake time beyond timestamp.
// func GetCmdAccountStakedPastTime() (*osmocli.QueryDescriptor, *types.AccountStakedPastTimeRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-staked-pastime <address> <timestamp>",
// 		Short: "Query staked records of an account with unstake time beyond timestamp",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-pastime <address> <timestamp>
// `}, &types.AccountStakedPastTimeRequest{}
// }

// // GetCmdAccountStakedPastTimeNotUnstakingOnly returns stakes of an account with unstake time beyond provided timestamp
// // amongst the stakes that are in the unstaking queue.
// func GetCmdAccountStakedPastTimeNotUnstakingOnly() (*osmocli.QueryDescriptor, *types.AccountStakedPastTimeNotUnstakingOnlyRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-staked-pastime-not-unstaking <address> <timestamp>",
// 		Short: "Query staked records of an account with unstake time beyond timestamp within not unstaking queue.",
// 		Long: `{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-pastime-not-unstaking <address> <timestamp>
// `}, &types.AccountStakedPastTimeNotUnstakingOnlyRequest{}
// }

// // GetCmdAccountUnstakedBeforeTime returns stakes with unstake time before the provided timestamp.
// func GetCmdAccountUnstakedBeforeTime() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountUnstakedBeforeTimeRequest](
// 		"account-staked-beforetime <address> <timestamp>",
// 		"Query account's unstaked records before specific time",
// 		`{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-pastime <address> <timestamp>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedPastTimeDenom returns stakes of an account whose unstake time is
// // beyond given timestamp, and stakes with the specified denom.
// func GetCmdAccountStakedPastTimeDenom() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedPastTimeDenomRequest](
// 		"account-staked-pastime-denom <address> <timestamp> <denom>",
// 		"Query account's stake records by address, timestamp, denom",
// 		`{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-pastime-denom <address> <timestamp> <denom>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdStakedByID returns stake by id.
// func GetCmdStakedByID() *cobra.Command {
// 	q := osmocli.QueryDescriptor{
// 		Use:   "stake-by-id <id>",
// 		Short: "Query account's stake record by id",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} stake-by-id 1`,
// 		QueryFnName: "StakedByID",
// 	}
// 	q.Long = osmocli.FormatLongDesc(q.Long, osmocli.NewLongMetadata(types.ModuleName).WithShort(q.Short))
// 	return osmocli.BuildQueryCli[*types.StakedRequest](&q, types.NewQueryClient)
// }

// // GetCmdNextStakeID returns next stake id to be created.
// func GetCmdNextStakeID() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.NextStakeIDRequest](
// 		"next-stake-id",
// 		"Query next stake id to be created",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedLongerDuration returns account staked records with longer duration.
// func GetCmdAccountStakedLongerDuration() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedLongerDurationRequest](
// 		"account-staked-longer-duration <address> <duration>",
// 		"Query account staked records with longer duration",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedLongerDuration returns account staked records with longer duration.
// func GetCmdAccountStakedDuration() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedDurationRequest](
// 		"account-staked-duration <address> <duration>",
// 		"Query account staked records with a specific duration",
// 		`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-staked-duration osmo1yl6hdjhmkf37639730gffanpzndzdpmhxy9ep3 604800s`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedLongerDurationNotUnstakingOnly returns account staked records with longer duration from unstaking only queue.
// func GetCmdAccountStakedLongerDurationNotUnstakingOnly() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedLongerDurationNotUnstakingOnlyRequest](
// 		"account-staked-longer-duration-not-unstaking <address> <duration>",
// 		"Query account staked records with longer duration from unstaking only queue",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountStakedLongerDurationDenom returns account's stakes for a specific denom
// // with longer duration than the given duration.
// func GetCmdAccountStakedLongerDurationDenom() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountStakedLongerDurationDenomRequest](
// 		"account-staked-longer-duration-denom <address> <duration> <denom>",
// 		"Query staked records for a denom with longer duration",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// func GetCmdTotalStakedByDenom() (*osmocli.QueryDescriptor, *types.StakedDenomRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "total-staked-of-denom <denom>",
// 		Short: "Query staked amount for a specific denom bigger then duration provided",
// 		Long: osmocli.FormatLongDescDirect(`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} total-staked-of-denom uosmo --min-duration=0s`, types.ModuleName),
// 	}, &types.StakedDenomRequest{}
// }

// // GetCmdOutputStakesJson outputs all stakes into a file called stake_export.json.
// func GetCmdOutputStakesJson() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "output-all-stakes <max stake ID>",
// 		Short: "output all stakes into a json file",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Output all stakes into a json file.
// Example:
// $ %s query stakeup output-all-stakes <max stake ID>
// `,
// 				version.AppName,
// 			),
// 		),
// 		Args: cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			maxStakeID, err := strconv.ParseInt(args[0], 10, 32)
// 			if err != nil {
// 				return err
// 			}

// 			// status
// 			const (
// 				doesnt_exist_status = iota
// 				unbonding_status
// 				bonded_status
// 			)

// 			type StakeResult struct {
// 				Id            int
// 				Status        int // one of {doesnt_exist, }
// 				Denom         string
// 				Amount        sdk.Int
// 				Address       string
// 				UnbondEndTime time.Time
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)

// 			results := []StakeResult{}
// 			for i := 0; i <= int(maxStakeID); i++ {
// 				curStakeResult := StakeResult{Id: i}
// 				res, err := queryClient.StakedByID(cmd.Context(), &types.StakedRequest{StakeId: uint64(i)})
// 				if err != nil {
// 					curStakeResult.Status = doesnt_exist_status
// 					results = append(results, curStakeResult)
// 					continue
// 				}
// 				// 1527019420 is hardcoded time well before launch, but well after year 1
// 				if res.Stake.EndTime.Before(time.Unix(1527019420, 0)) {
// 					curStakeResult.Status = bonded_status
// 				} else {
// 					curStakeResult.Status = unbonding_status
// 					curStakeResult.UnbondEndTime = res.Stake.EndTime
// 					curStakeResult.Denom = res.Stake.Coins[0].Denom
// 					curStakeResult.Amount = res.Stake.Coins[0].Amount
// 					curStakeResult.Address = res.Stake.Owner
// 				}
// 				results = append(results, curStakeResult)
// 			}

// 			bz, err := json.Marshal(results)
// 			if err != nil {
// 				return err
// 			}
// 			err = os.WriteFile("stake_export.json", bz, 0o777)
// 			if err != nil {
// 				return err
// 			}

// 			fmt.Println("Writing to stake_export.json")
// 			return nil
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }
