// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a tree backed by a red-black tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package treeset

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections"
	rbt "github.com/jabong/florest-core/src/common/collections/trees/rbtree"
	"strings"
)

// Set holds elements in a red-black tree
type Set struct {
	tree *rbt.Tree
}

var itemExists = struct{}{}

// NewWith instantiates a new empty set with the custom comparator.
func NewWith(comparator collections.Comparator) *Set {
	return &Set{tree: rbt.New(comparator)}
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		_, found := set.tree.Get(item)
		if !found {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	return set.tree.Keys()
}

// Iterator holding the iterator's state
func (set *Set) Iterator() collections.Iterator {
	return &Iterator{index: 0, rbIterator: set.tree.Iterator()}
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "TreeSet\n"
	for _, v := range set.Values() {
		str += fmt.Sprintf("%v,", v)
	}
	return strings.TrimRight(str, ",")
}
