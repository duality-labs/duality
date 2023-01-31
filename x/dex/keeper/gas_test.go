package keeper_test

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) DoRun(nTicks int, step int, withLo bool) Data {
	fmt.Printf("Start run with params %v, %v\n", nTicks, step)

	s.fundAliceBalances(1000000, 1000000)
	s.fundBobBalances(1000000, 1000000)
	i := 0
	amountDeposited := 0
	for ; i < nTicks; i++ {
		curIdx := i * step
		s.aliceDeposits(
			NewDeposit(0, 5, curIdx, 0),
			NewDeposit(0, 5, curIdx, 1))

		amountDeposited += 10
		if withLo {
			s.aliceLimitSells("TokenB", curIdx, 5)
			amountDeposited += 5
		}
	}
	startTime := time.Now()
	startGas := s.ctx.GasMeter().GasConsumed()
	amountToSwap := CalcAmountToSwap((nTicks-1)*step, amountDeposited)
	fmt.Printf("Amount int: %v, amountToSwap: %v\n", amountDeposited, amountToSwap)

	s.bobMarketSells("TokenA", int(amountToSwap), 0)
	duration := time.Since(startTime)
	endGas := s.ctx.GasMeter().GasConsumed()
	totalGas := endGas - startGas
	return Data{
		NTicks:   nTicks,
		Step:     step,
		GasUsed:  totalGas,
		Duration: int64(duration),
	}
}

func WriteCSV(data []Data) {
	f, err := os.Create("/Users/julian/Desktop/gas_test_data.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, d := range data {
		if err := w.Write(d.StringArr()); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func (s *MsgServerTestSuite) TestGasSwap() {
	var data []Data
	tickAmts := []int{1, 2, 10, 20, 100, 200, 1000}
	steps := []int{1, 2, 5, 10, 50}
	for _, tickAmt := range tickAmts {
		for _, step := range steps {
			s.SetupTest()
			runData := s.DoRun(tickAmt, step, false)
			data = append(data, runData)
		}
	}

	WriteCSV(data)
	fmt.Printf("All Data: %v\n", data)

}

type Data struct {
	NTicks   int
	Step     int
	GasUsed  sdk.Gas
	Duration int64
}

func (d Data) StringArr() []string {
	return []string{
		strconv.Itoa(d.NTicks),
		strconv.Itoa(d.Step),
		strconv.FormatUint(d.GasUsed, 10),
		strconv.FormatInt(d.Duration, 10)}
}

func CalcAmountToSwap(maxtick int, amountIn int) int64 {
	worstPrice := sdk.OneDec().Quo(Pow(BasePrice(), uint64(maxtick)))
	return sdk.NewDec(int64(amountIn)).Quo(worstPrice).Ceil().TruncateInt64()
}

func BasePrice() sdk.Dec {
	return sdk.MustNewDecFromStr("1.0001")
}

func Pow(a sdk.Dec, n uint64) sdk.Dec {
	if n == 0 {
		return sdk.OneDec()
	}
	if n&1 == 0 {
		return Pow(a.Mul(a), n>>1)
	} else {
		return a.Mul(Pow(a.Mul(a), n>>1))
	}
}
