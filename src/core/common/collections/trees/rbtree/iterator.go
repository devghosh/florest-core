// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import "github.com/emirpasic/gods/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	tree       *Tree
	node       *Node
	state      state
	nextCalled bool
	hasNext    bool
}

type state byte

const (
	begin, between, end state = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree) Iterator() Iterator {
	return Iterator{tree: Tree, node: nil, state: begin, nextCalled: true}
}

func (iterator *Iterator) HasNext() bool {
	if !iterator.nextCalled {
		return iterator.hasNext
	}
	if iterator.state == end {
		return iterator.goToEnd()
	}
	if iterator.state == begin {
		left := iterator.tree.Left()
		if left == nil {
			return iterator.goToEnd()
		}
		iterator.node = left
		return iterator.goInBetween()
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		return iterator.goInBetween()
	}
	node := iterator.node
	for iterator.node.Parent != nil {
		iterator.node = iterator.node.Parent
		if iterator.tree.Comparator.Compare(node.Key, iterator.node.Key) <= 0 {
			return iterator.goInBetween()
		}
	}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() *Entry {
	if iterator.HasNext() {
		iterator.nextCalled = true
		return NewEntry(iterator.node.Key, iterator.node.Value)
	}
	return nil
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.node = nil
	iterator.state = begin
	iterator.nextCalled = true
}

func (iterator *Iterator) goInBetween() bool {
	iterator.state = between
	iterator.hasNext = true
	iterator.nextCalled = false
	return true
}

func (iterator *Iterator) goToEnd() bool {
	iterator.node = nil
	iterator.state = end
	iterator.hasNext = false
	return false
}
