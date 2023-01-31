package keeper_test

import "fmt"

const (
	step  = 1
	start = 3
	end   = 200
)

func (s *MsgServerTestSuite) TestGasSwap() {
	s.fundAliceBalances(1000000, 1000000)
	s.fundBobBalances(1000000, 1000000)
	curIdx := start
	amountDeposited := 0
	for ; curIdx < end; curIdx++ {
		s.aliceDeposits(
			NewDeposit(5, 0, curIdx*-1, 0),
			NewDeposit(0, 5, curIdx, 0))
		amountDeposited += 5
	}
	startGas := s.ctx.GasMeter().GasConsumed()
	s.bobMarketSells("TokenB", amountDeposited, 0)
	endGas := s.ctx.GasMeter().GasConsumed()
	totalGas := endGas - startGas
	fmt.Printf("Gas used: %v\n", totalGas)

}
