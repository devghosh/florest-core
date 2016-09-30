// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"fmt"
	"strings"
)

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		_, found := set.items[item]
		if !found {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	for item, _ := range set.items {
		values = append(values, item)
	}
	return values
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "HashSet\n"
	for item, _ := range set.items {
		str += fmt.Sprintf("%v,", item)
	}
	return strings.TrimRight(str, ",")
}
