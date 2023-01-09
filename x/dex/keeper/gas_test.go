// TEMP: Just for testing speed differences. Will remove
package keeper_test

import "fmt"

const nTicks int = 10
const spacing = 2

func (s *MsgServerTestSuite) TestTickGas() {

	var i int = 0
	for i > nTicks*-1 {
		tick, _ := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", int64(i))
		i = i - spacing
		s.app.DexKeeper.SetTick(s.ctx, "TokenA<>TokenB", tick)
	}
	s.fundAliceBalances(20, 0)
	s.fundBobBalances(0, 20)
	s.aliceDeposits(
		NewDeposit(1, 0, 0, 0),
		NewDeposit(1, 0, -11, 0),
		NewDeposit(1, 0, -12, 0),
		NewDeposit(1, 0, -13, 0),
		NewDeposit(1, 0, -14, 0),
		NewDeposit(5, 0, -15, 0),
		NewDeposit(10, 0, -1*nTicks, 0),
	)
	SpentUntilNow := s.ctx.GasMeter().GasConsumed()

	s.bobMarketSells("TokenB", 20, 17)
	SpentFinal := s.ctx.GasMeter().GasConsumed()
	SpentTotal := SpentFinal - SpentUntilNow
	fmt.Printf("Gas Spent: %v\n", SpentTotal)

}

func (s *MsgServerTestSuite) TestTickGas2() {

	var i int = 0
	for i < nTicks {
		tick, _ := s.app.DexKeeper.GetOrInitTick(s.goCtx, "TokenA<>TokenB", int64(i))
		i = i + spacing
		s.app.DexKeeper.SetTick(s.ctx, "TokenA<>TokenB", tick)
	}
	s.fundAliceBalances(0, 20)
	s.fundBobBalances(20, 0)
	s.aliceDeposits(
		NewDeposit(0, 1, 0, 0),
		NewDeposit(0, 1, 11, 0),
		NewDeposit(0, 1, 12, 0),
		NewDeposit(0, 1, 13, 0),
		NewDeposit(0, 1, 14, 0),
		NewDeposit(0, 5, 15, 0),
		NewDeposit(0, 10, nTicks, 0),
	)
	SpentUntilNow := s.ctx.GasMeter().GasConsumed()

	s.bobMarketSells("TokenA", 20, 17)
	SpentFinal := s.ctx.GasMeter().GasConsumed()
	SpentTotal := SpentFinal - SpentUntilNow
	fmt.Printf("Gas Spent: %v\n", SpentTotal)

}
