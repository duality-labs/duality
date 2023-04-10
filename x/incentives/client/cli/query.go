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
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdModuleLockedAmount)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountUnlockingCoins)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountLockedPastTime)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdAccountLockedPastTimeNotUnlockingOnly)
	// osmocli.AddQueryCmd(cmd, qcGetter, GetCmdTotalLockedByDenom)
	cmd.AddCommand(
	// GetCmdRewardsEstimate(),
	// GetCmdAccountUnlockableCoins(),
	// GetCmdAccountLockedCoins(),
	// GetCmdAccountUnlockedBeforeTime(),
	// GetCmdAccountLockedPastTimeDenom(),
	// GetCmdLockedByID(),
	// GetCmdAccountLockedLongerDuration(),
	// GetCmdAccountLockedLongerDurationNotUnlockingOnly(),
	// GetCmdAccountLockedLongerDurationDenom(),
	// GetCmdOutputLocksJson(),
	// GetCmdAccountLockedDuration(),
	// GetCmdNextLockID(),
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
// 			var res *types.AccountLockedLongerDurationResponse
// 			ownerLocks := []uint64{}
// 			lockIds := []uint64{}

// 			owner, err := cmd.Flags().GetString(FlagOwner)
// 			if err != nil {
// 				return err
// 			}

// 			lockIdsCombined, err := cmd.Flags().GetString(FlagLockIds)
// 			if err != nil {
// 				return err
// 			}
// 			lockIdStrs := strings.Split(lockIdsCombined, ",")

// 			endEpoch, err := cmd.Flags().GetInt64(FlagEndEpoch)
// 			if err != nil {
// 				return err
// 			}

// 			// if user doesn't provide at least one of the lock ids or owner, we don't have enough information to proceed.
// 			if lockIdsCombined == "" && owner == "" {
// 				return fmt.Errorf("either one of owner flag or lock IDs must be provided")

// 				// if user provides lockIDs, use these lockIDs in our rewards estimation
// 			} else if owner == "" {
// 				for _, lockIdStr := range lockIdStrs {
// 					lockId, err := strconv.ParseUint(lockIdStr, 10, 64)
// 					if err != nil {
// 						return err
// 					}
// 					lockIds = append(lockIds, lockId)
// 				}

// 				// if no lockIDs are provided but an owner is provided, we query the rewards for all of the locks the owner has
// 			} else if lockIdsCombined == "" {
// 				lockIds = append(lockIds, ownerLocks...)
// 			}

// 			// if lockIDs are provided and an owner is provided, only query the lockIDs that are provided
// 			// if a lockID was provided and it doesn't belong to the owner, return an error
// 			if len(lockIds) != 0 && owner != "" {
// 				for _, lockId := range lockIds {
// 					validInputLockId := contains(ownerLocks, lockId)
// 					if !validInputLockId {
// 						return fmt.Errorf("lock-id %v does not belong to %v", lockId, owner)
// 					}
// 				}
// 			}

// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)

// 			if owner != "" {
// 				queryClientLockup := types.NewQueryClient(clientCtx)

// 				res, err = queryClientLockup.AccountLockedLongerDuration(cmd.Context(), &types.AccountLockedLongerDurationRequest{Owner: owner, Duration: time.Millisecond})
// 				if err != nil {
// 					return err
// 				}
// 				for _, lockId := range res.Locks {
// 					ownerLocks = append(ownerLocks, lockId.ID)
// 				}
// 			}

// 			// TODO: Fix accumulation store bug. For now, we return a graceful error when attempting to query bugged gauges
// 			RewardsEstimateResult, err := queryClient.RewardsEstimate(cmd.Context(), &types.RewardsEstimateRequest{
// 				Owner:    owner, // owner is used only when lockIds are empty
// 				LockIds:  lockIds,
// 				EndEpoch: endEpoch,
// 			})
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(RewardsEstimateResult)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)
// 	cmd.Flags().String(FlagOwner, "", "Owner to receive rewards, optionally used when lock-ids flag is NOT set")
// 	cmd.Flags().String(FlagLockIds, "", "the lock ids to receive rewards, when it is empty, all lock ids of the owner are used")
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

// // GetCmdModuleBalance returns full balance of the lockup module.
// // Lockup module is where coins of locks are held.
// // This includes locked balance and unlocked balance of the module.
// func GetCmdModuleBalance() (*osmocli.QueryDescriptor, *types.ModuleBalanceRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "module-balance",
// 		Short: "Query module balance",
// 		Long:  `{{.Short}}`}, &types.ModuleBalanceRequest{}
// }

// // GetCmdModuleLockedAmount returns locked balance of the module,
// // which are all the tokens not unlocking + tokens that are not finished unlocking.
// func GetCmdModuleLockedAmount() (*osmocli.QueryDescriptor, *types.ModuleLockedAmountRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "module-locked-amount",
// 		Short: "Query locked amount",
// 		Long:  `{{.Short}}`}, &types.ModuleLockedAmountRequest{}
// }

// // GetCmdAccountUnlockableCoins returns unlockable coins which has finsihed unlocking.
// // TODO: DELETE THIS + Actual query in subsequent PR
// func GetCmdAccountUnlockableCoins() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "account-unlockable-coins <address>",
// 		Short: "Query account's unlockable coins",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query account's unlockable coins.

// Example:
// $ %s query lockup account-unlockable-coins <address>
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

// 			res, err := queryClient.AccountUnlockableCoins(cmd.Context(), &types.AccountUnlockableCoinsRequest{Owner: args[0]})
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(res)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }

// // GetCmdAccountUnlockingCoins returns unlocking coins of a specific account.
// func GetCmdAccountUnlockingCoins() (*osmocli.QueryDescriptor, *types.AccountUnlockingCoinsRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-unlocking-coins <address>",
// 		Short: "Query account's unlocking coins",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-unlocking-coins <address>`}, &types.AccountUnlockingCoinsRequest{}
// }

// // GetCmdAccountLockedCoins returns locked coins that that are still in a locked state from the specified account.
// func GetCmdAccountLockedCoins() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedCoinsRequest](
// 		"account-locked-coins <address>",
// 		"Query account's locked coins",
// 		`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-coins <address>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedPastTime returns locks of an account with unlock time beyond timestamp.
// func GetCmdAccountLockedPastTime() (*osmocli.QueryDescriptor, *types.AccountLockedPastTimeRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-locked-pastime <address> <timestamp>",
// 		Short: "Query locked records of an account with unlock time beyond timestamp",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-pastime <address> <timestamp>
// `}, &types.AccountLockedPastTimeRequest{}
// }

// // GetCmdAccountLockedPastTimeNotUnlockingOnly returns locks of an account with unlock time beyond provided timestamp
// // amongst the locks that are in the unlocking queue.
// func GetCmdAccountLockedPastTimeNotUnlockingOnly() (*osmocli.QueryDescriptor, *types.AccountLockedPastTimeNotUnlockingOnlyRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "account-locked-pastime-not-unlocking <address> <timestamp>",
// 		Short: "Query locked records of an account with unlock time beyond timestamp within not unlocking queue.",
// 		Long: `{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-pastime-not-unlocking <address> <timestamp>
// `}, &types.AccountLockedPastTimeNotUnlockingOnlyRequest{}
// }

// // GetCmdAccountUnlockedBeforeTime returns locks with unlock time before the provided timestamp.
// func GetCmdAccountUnlockedBeforeTime() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountUnlockedBeforeTimeRequest](
// 		"account-locked-beforetime <address> <timestamp>",
// 		"Query account's unlocked records before specific time",
// 		`{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-pastime <address> <timestamp>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedPastTimeDenom returns locks of an account whose unlock time is
// // beyond given timestamp, and locks with the specified denom.
// func GetCmdAccountLockedPastTimeDenom() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedPastTimeDenomRequest](
// 		"account-locked-pastime-denom <address> <timestamp> <denom>",
// 		"Query account's lock records by address, timestamp, denom",
// 		`{{.Short}}
// Timestamp is UNIX time in seconds.{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-pastime-denom <address> <timestamp> <denom>
// `, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdLockedByID returns lock by id.
// func GetCmdLockedByID() *cobra.Command {
// 	q := osmocli.QueryDescriptor{
// 		Use:   "lock-by-id <id>",
// 		Short: "Query account's lock record by id",
// 		Long: `{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} lock-by-id 1`,
// 		QueryFnName: "LockedByID",
// 	}
// 	q.Long = osmocli.FormatLongDesc(q.Long, osmocli.NewLongMetadata(types.ModuleName).WithShort(q.Short))
// 	return osmocli.BuildQueryCli[*types.LockedRequest](&q, types.NewQueryClient)
// }

// // GetCmdNextLockID returns next lock id to be created.
// func GetCmdNextLockID() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.NextLockIDRequest](
// 		"next-lock-id",
// 		"Query next lock id to be created",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedLongerDuration returns account locked records with longer duration.
// func GetCmdAccountLockedLongerDuration() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationRequest](
// 		"account-locked-longer-duration <address> <duration>",
// 		"Query account locked records with longer duration",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedLongerDuration returns account locked records with longer duration.
// func GetCmdAccountLockedDuration() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedDurationRequest](
// 		"account-locked-duration <address> <duration>",
// 		"Query account locked records with a specific duration",
// 		`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} account-locked-duration osmo1yl6hdjhmkf37639730gffanpzndzdpmhxy9ep3 604800s`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedLongerDurationNotUnlockingOnly returns account locked records with longer duration from unlocking only queue.
// func GetCmdAccountLockedLongerDurationNotUnlockingOnly() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationNotUnlockingOnlyRequest](
// 		"account-locked-longer-duration-not-unlocking <address> <duration>",
// 		"Query account locked records with longer duration from unlocking only queue",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// // GetCmdAccountLockedLongerDurationDenom returns account's locks for a specific denom
// // with longer duration than the given duration.
// func GetCmdAccountLockedLongerDurationDenom() *cobra.Command {
// 	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationDenomRequest](
// 		"account-locked-longer-duration-denom <address> <duration> <denom>",
// 		"Query locked records for a denom with longer duration",
// 		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
// }

// func GetCmdTotalLockedByDenom() (*osmocli.QueryDescriptor, *types.LockedDenomRequest) {
// 	return &osmocli.QueryDescriptor{
// 		Use:   "total-locked-of-denom <denom>",
// 		Short: "Query locked amount for a specific denom bigger then duration provided",
// 		Long: osmocli.FormatLongDescDirect(`{{.Short}}{{.ExampleHeader}}
// {{.CommandPrefix}} total-locked-of-denom uosmo --min-duration=0s`, types.ModuleName),
// 	}, &types.LockedDenomRequest{}
// }

// // GetCmdOutputLocksJson outputs all locks into a file called lock_export.json.
// func GetCmdOutputLocksJson() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "output-all-locks <max lock ID>",
// 		Short: "output all locks into a json file",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Output all locks into a json file.
// Example:
// $ %s query lockup output-all-locks <max lock ID>
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

// 			maxLockID, err := strconv.ParseInt(args[0], 10, 32)
// 			if err != nil {
// 				return err
// 			}

// 			// status
// 			const (
// 				doesnt_exist_status = iota
// 				unbonding_status
// 				bonded_status
// 			)

// 			type LockResult struct {
// 				Id            int
// 				Status        int // one of {doesnt_exist, }
// 				Denom         string
// 				Amount        sdk.Int
// 				Address       string
// 				UnbondEndTime time.Time
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)

// 			results := []LockResult{}
// 			for i := 0; i <= int(maxLockID); i++ {
// 				curLockResult := LockResult{Id: i}
// 				res, err := queryClient.LockedByID(cmd.Context(), &types.LockedRequest{LockId: uint64(i)})
// 				if err != nil {
// 					curLockResult.Status = doesnt_exist_status
// 					results = append(results, curLockResult)
// 					continue
// 				}
// 				// 1527019420 is hardcoded time well before launch, but well after year 1
// 				if res.Lock.EndTime.Before(time.Unix(1527019420, 0)) {
// 					curLockResult.Status = bonded_status
// 				} else {
// 					curLockResult.Status = unbonding_status
// 					curLockResult.UnbondEndTime = res.Lock.EndTime
// 					curLockResult.Denom = res.Lock.Coins[0].Denom
// 					curLockResult.Amount = res.Lock.Coins[0].Amount
// 					curLockResult.Address = res.Lock.Owner
// 				}
// 				results = append(results, curLockResult)
// 			}

// 			bz, err := json.Marshal(results)
// 			if err != nil {
// 				return err
// 			}
// 			err = os.WriteFile("lock_export.json", bz, 0o777)
// 			if err != nil {
// 				return err
// 			}

// 			fmt.Println("Writing to lock_export.json")
// 			return nil
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }
