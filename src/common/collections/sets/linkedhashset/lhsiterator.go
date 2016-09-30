// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator holding the iterator's state
type Iterator struct {
	iterator collections.Iterator
	index    int
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) HasNext() bool {
	return iterator.iterator.HasNext()
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.iterator.Next()
	if temp == nil {
		return nil
	}
	index := iterator.index
	iterator.index++
	return collections.NewEntry(index, temp.GetKey())
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.iterator.Reset()
	iterator.index = 0
}
