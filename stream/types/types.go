package types

type (
	T interface{}
	R interface{}

	Predicate   func(T) bool
	Function    func(T) R
	Consumer    func(T)
	Comparator  func(T, T) int
	Accumulator func(R, T) R
)
