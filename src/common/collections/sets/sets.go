package sets

import "github.com/jabong/florest-core/src/common/collections"

// Set interface that all sets implement
type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})

	collections.Collection
}
