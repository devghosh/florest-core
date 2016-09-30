package maps

import "github.com/jabong/florest-core/src/common/collections"

// Map interface that all maps implement
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}

	collections.Collection
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// Contains(keys ...interface{}) bool
}
