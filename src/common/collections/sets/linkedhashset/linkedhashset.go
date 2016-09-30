package linkedhashset

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections"
	lhmap "github.com/jabong/florest-core/src/common/collections/maps/linkedhashmap"
	"strings"
)

type Set struct {
	items *lhmap.Map
}

var itemExists = struct{}{}

// New instantiates a hash map.
func New() *Set {
	return &Set{items: lhmap.New()}
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		set.items.Remove(item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	return set.items.Contains(items)
}

// Empty returns true if map does not contain any elements
func (set *Set) IsEmpty() bool {
	return set.items.IsEmpty()
}

// Size returns number of elements in the map.
func (set *Set) Size() int {
	return set.items.Size()
}

// Values returns all values (insertion order).
func (set *Set) Values() []interface{} {
	return set.items.Keys()
}

// Clear removes all elements from the map.
func (set *Set) Clear() {
	set.items.Clear()
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (set *Set) Iterator() collections.Iterator {
	return &Iterator{iterator: set.items.Iterator(), index: 0}
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "LinkedHashSet\n"
	for _, v := range set.Values() {
		str += fmt.Sprintf("%v,", v)
	}
	return strings.TrimRight(str, ",")
}
