package keeper

import (
	"context"
	//"fmt"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


// Specifications of types.Msg.SingleDeposit can be found in ../proto/dex/tx.proto

func (k msgServer) SingleDeposit(goCtx context.Context, msg *types.MsgSingleDeposit) (*types.MsgSingleDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	/* Note: while we want to intially check the validity of the sdk.AccAddress receiver address, 
	the string version: msg.Receiver is used in our send call
	*/
	_ = receiverAddr

	
	//Coverts msg.Amounts0 (string) to sdk.Dec
	// Note: msg.Amounts0 should be specified as a decimal as input (ie "1.45")
	amount0, err := sdk.NewDecFromStr(msg.Amounts0)
	// Error checking for valid sdk.Dec
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	//Coverts msg.Amounts1 (string) to sdk.Dec
	amount1, err := sdk.NewDecFromStr(msg.Amounts1)
	// Error checking for valid sdk.De
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	/* GetBalance is an external (x/bank) module function called to check the token balanced for a specified address
	  GetBalance returns a sdk.Coin: (string denom sdk.Int amount)

	  Note: sdk.Int is not the same as native Int but is a override of big/int

	  AccountsToken0Balance is the sdk.Dec representation of the return sdk.Coin's amount of token 0.
	*/

	AccountsToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token0).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsToken0Balance.LT(amount0) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	//  AccountsToken1Balance is the sdk.Dec representation of the return sdk.Coin's amount of token 1.
	AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token1).Amount)
	if AccountsToken1Balance.LT(amount1) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s does not have enough  of token 1", callerAddr)
	}

	// Sorts token0, token1 for uniformity for use in internal mappings, if sorting is needed amounts0, amounts1 are switched as well.
	token0, token1, amounts0, amounts1, error := k.SortTokensDeposit(ctx, msg.Token0, msg.Token1, []sdk.Dec{amount0}, []sdk.Dec{amount1})

	// Error handling for SortTokensDeposit
	if error != nil {
		return nil, error
	}

	// In the case of a single deposit we have amounts0, amounts1 will be arrays of length 1.
	amount0 = amounts0[0]
	amount1 = amounts1[0]
	
	// Determines if previous shares exists for a address at a specified token pair, price, fee.
	shareOld, shareFound := k.GetShare(
		ctx,
		msg.Receiver,
		token0,
		token1,
		msg.Price,
		msg.Fee,
	)

	if !shareFound {
		shareOld = types.Share{
			Owner:       msg.Receiver,
			Token0:      token0,
			Token1:      token1,
			Price:       msg.Price,
			Fee:         msg.Fee,
			ShareAmount: sdk.ZeroDec(),
		}
	}

	//fmt.Println("All ticks contracts:", k.GetAllTicks(ctx))

	// Determines if a given token pair has been previously initialized
	tickOld, tickFound := k.GetTicks(
		ctx,
		token0,
		token1,
	)

	// Converts msg.Price (string) to a sdk.Dec
	price, err := sdk.NewDecFromStr(msg.Price)

	// Error handling for valid sdk.Dec
	if err != nil {
		return nil, err
	}

	// Converts msg.Fee (string) to a sdk.Dec
	fee, err := sdk.NewDecFromStr(msg.Fee)

	// Error handling for a valid sdk.Dec
	if err != nil {
		return nil, err
	}

	// Internal variable initializations
	var OneToZeroOld types.Pool
	var ZeroToOneOld types.Pool
	OneToZeroFound := false
	ZeroToOneFound := false

	var SharesMinted sdk.Dec
	var trueAmounts0 = amount0
	var trueAmounts1 = amount1

	
	if tickFound {

		// Check to see which pool's previously exist
		OneToZeroOld, OneToZeroFound = k.GetPool(&tickOld.PoolsOneToZero, fee, price)
		ZeroToOneOld, ZeroToOneFound = k.GetPool(&tickOld.PoolsZeroToOne, fee, price)

		// If either OneToZero or ZeroToOne is found we calculate the correct amount of amounts0, amounts1 to deposit by calling  DepositHelperAdd
		if OneToZeroFound {
			trueAmounts0, trueAmounts1, SharesMinted, err = k.DepositHelperAdd(&OneToZeroOld, amount0, amount1)

			if err != nil {
				return nil, err
			}
		} else if ZeroToOneFound {
			trueAmounts0, trueAmounts1, SharesMinted, err = k.DepositHelperAdd(&ZeroToOneOld, amount0, amount1)

			if err != nil {
				return nil, err
			}

			// Neither pool has been found but  the tick has been previously initialized, caluclate sharesAmounts, and trueAmounts as if it
			// is newly being initialized.
		} else if !OneToZeroFound && !ZeroToOneFound {

			SharesMinted = amount0.Add(amount1.Mul(price))
		
		}
	
		// No token pair was found, specified pool is initialized. 
	} else {
		SharesMinted = amount0.Add(amount1.Mul(price))
	}

	var NewPool types.Pool

	// Sets the updated pool (NewPool) via either updating previous pool information, with respect to what pools pools previously existsed
	// (ie if OneToZero exists we calculate NewPool from its state)
	if OneToZeroFound {
		NewPool = types.Pool{
			Reserve0:    OneToZeroOld.Reserve0.Add(trueAmounts0),
			Reserve1:    OneToZeroOld.Reserve1.Add(trueAmounts1),
			Fee:         fee,
			Price:       price,
			TotalShares: OneToZeroOld.TotalShares.Add(SharesMinted),
			Index:       0,
		}
	} else if ZeroToOneFound {
		NewPool = types.Pool{
			Reserve0:    ZeroToOneOld.Reserve0.Add(trueAmounts0),
			Reserve1:    ZeroToOneOld.Reserve1.Add(trueAmounts1),
			Fee:         fee,
			Price:       price,
			TotalShares: ZeroToOneOld.TotalShares.Add(SharesMinted),
			Index:       0,
		}

	// If no pool is found, we initialize a New Pool based off of our calculations above
	} else {
		NewPool = types.Pool{
			Reserve0:    trueAmounts0,
			Reserve1:    trueAmounts1,
			Fee:         fee,
			Price:       price,
			TotalShares: SharesMinted,
			Index:       0,
		}
	}

	//Sends TrueAmounts0 from msg.Sender to module
	if trueAmounts0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(trueAmounts0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}

	//Sends TrueAmounts1 from msg.Sender to module
	if trueAmounts1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(trueAmounts1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}

	// If the ZeroToOne pool exists we update the priority queue based off of updates of NewPool to the previous state
	if ZeroToOneFound {
		k.Update0to1(&tickOld.PoolsZeroToOne, &ZeroToOneOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)

	// Pushes a NewPool to the priority queue given that reserves1 is now greater than 0, and that this pool did not previously exists in the priority queue
	} else if NewPool.Reserve1.GT(sdk.ZeroDec()) && !ZeroToOneFound {
		
		k.Push0to1(&tickOld.PoolsZeroToOne, &NewPool)
	}

	// If the OneToZero pool exists we update the priority queue based off of updates of NewPool to the previous state
	if OneToZeroFound {
		
		k.Update1to0(&tickOld.PoolsOneToZero, &OneToZeroOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)

	// Pushes a NewPool to the priority queue given that reserves0 is now greater than 0, and that this pool did not previously exists in the priority queue
	} else if NewPool.Reserve0.GT(sdk.ZeroDec()) && !OneToZeroFound {
		k.Push1to0(&tickOld.PoolsOneToZero, &NewPool)
	}

	// Initialize a newTick object based off of our updates via deposits / edits to priority queues.
	// Note: all edits of the priority queues are done by reference and thus we only need to pass in the original objects from k.GetTicks
	tickNew := types.Ticks{
		Token0:         token0,
		Token1:         token1,
		PoolsZeroToOne: tickOld.PoolsZeroToOne,
		PoolsOneToZero: tickOld.PoolsOneToZero,
	}

	// Initialize share Object based of deposit updates
	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0,
		Token1:      token1,
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: shareOld.ShareAmount.Add(SharesMinted),
	}

	k.SetTicks(
		ctx,
		tickNew,
	)

	k.SetShare(
		ctx,
		shareNew,
	)

	var event = sdk.NewEvent(sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.DepositEventKey),
		sdk.NewAttribute(types.DepositEventCreator, msg.Creator),
		sdk.NewAttribute(types.DepositEventToken0, token0),
		sdk.NewAttribute(types.DepositEventToken1, token1),
		sdk.NewAttribute(types.DepositEventPrice, msg.Price),
		sdk.NewAttribute(types.DepositEventFee, msg.Fee),
		sdk.NewAttribute(types.DepositEventNewReserves0, NewPool.Reserve0.String()),
		sdk.NewAttribute(types.DepositEventNewReserves1, NewPool.Reserve1.String()),
		sdk.NewAttribute(types.DepositEventReceiver, msg.Receiver),
		sdk.NewAttribute(types.DepositEventSharesMinted, SharesMinted.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSingleDepositResponse{SharesMinted.String()}, nil
}
