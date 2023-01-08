package types

import "fmt"

func (p *PairId) Stringify() string {
	return fmt.Sprintf("%s<>%s", p.Token0, p.Token1)
}
