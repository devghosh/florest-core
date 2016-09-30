package collections

// Collection is base interface that all data structures implement.
type Collection interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []interface{}
	Contains(keys ...interface{}) bool
}
