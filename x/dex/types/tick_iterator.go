package types

type TickIteratorI interface {
	Next()
	Valid() bool
	Close() error
	Value() Tick
}
