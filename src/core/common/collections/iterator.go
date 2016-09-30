package collections

// Iterator is stateful iterator for ordered containers whose values can be fetched by an index.
type Iterator interface {
	HasNext() bool

	Next() *Entry

	// Reset method resets the iterator to its initial state
	Reset()
}
