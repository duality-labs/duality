package types

type TickIteratorI interface {
	Next() (int64, bool)
}
