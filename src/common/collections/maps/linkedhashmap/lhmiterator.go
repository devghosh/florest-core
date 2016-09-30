// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator holding the iterator's state
type Iterator struct {
	m       *Map
	current *Link
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) HasNext() bool {
	if iterator.current == nil {
		return false
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.current
	if temp == nil {
		return nil
	}
	iterator.current = temp.next
	return collections.NewEntry(temp.key, temp.value)
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.current = iterator.m.head
}
