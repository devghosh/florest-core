package sortedset

import (
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/sets"
)

type SortedSet interface {
	First() interface{}
	Last() interface{}
	sets.Set
	collections.Iterable
	collections.Comparable
}
