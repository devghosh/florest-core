package sortedmap

import (
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/maps"
)

type SortedMap interface {
	FirstKey() interface{}
	LastKey() interface{}
	maps.Map
	collections.Iterable
	collections.Comparable
}
