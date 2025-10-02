package bourke

import (
	"errors"
	"iter"
	"slices"
)

func (trie *trie[K, V]) Put(key []K, value V) {
	currentPrefix := trie.root
	endIndex := len(key) - 1
	for index := 0; index < endIndex; index++ {
		currentPrefix = trie.makeInternalPrefix(currentPrefix, key[index])
	}
	if currentPrefix.transitions == nil {
		transitions := newTree[K, *prefix[K, V]]()
		_ = transitions.Put(key[endIndex], &prefix[K, V]{
			parent:      currentPrefix,
			parentEdge:  key[endIndex],
			transitions: nil,
			flags:       prefixKey,
			value:       value,
		})
		currentPrefix.transitions = transitions
		trie.internalSize++
		trie.size++
		return
	}
	transition := currentPrefix.transitions.get(key[endIndex])
	if transition == nil {
		_ = currentPrefix.transitions.Put(key[endIndex], &prefix[K, V]{
			parent:      currentPrefix,
			parentEdge:  key[endIndex],
			transitions: nil,
			flags:       prefixKey,
			value:       value,
		})
		trie.internalSize++
		trie.size++
		return
	}
	if transition.value.isTombstone() {
		trie.size++
	}
	transition.value.parent = currentPrefix
	transition.value.parentEdge = key[endIndex]
	transition.value.flags = (transition.value.flags | prefixKey) & ^prefixTombstone
	transition.value.value = value
	_ = currentPrefix.transitions.Put(key[endIndex], transition.value)
}

func (trie *trie[K, V]) Tombstone(key []K) {
	currentPrefix := trie.root
	endIndex := len(key) - 1
	for index := 0; index < endIndex; index++ {
		currentPrefix = trie.makeInternalPrefix(currentPrefix, key[index])
	}
	if currentPrefix.transitions == nil {
		transitions := newTree[K, *prefix[K, V]]()
		_ = transitions.Put(key[endIndex], &prefix[K, V]{
			parent:      currentPrefix,
			parentEdge:  key[endIndex],
			transitions: nil,
			flags:       prefixTombstone,
		})
		currentPrefix.transitions = transitions
		trie.internalSize++
		return
	}
	transition := currentPrefix.transitions.get(key[endIndex])
	if transition == nil {
		_ = currentPrefix.transitions.Put(key[endIndex], &prefix[K, V]{
			parent:      currentPrefix,
			parentEdge:  key[endIndex],
			transitions: nil,
			flags:       prefixTombstone,
		})
		trie.internalSize++
		return
	}
	if transition.value.isKey() {
		trie.size--
	}
	transition.value.parent = currentPrefix
	transition.value.parentEdge = key[endIndex]
	transition.value.flags = (transition.value.flags | prefixTombstone) & ^prefixKey
	_ = currentPrefix.transitions.Put(key[endIndex], transition.value)
}

func (trie *trie[K, V]) Remove(key []K) {
	currentPrefix := trie.get(key)
	if (currentPrefix != nil) && currentPrefix.isKey() {
		for {
			if (currentPrefix.transitions == nil) && (currentPrefix.parent != nil) {
				_ = currentPrefix.parent.transitions.Remove(currentPrefix.parentEdge)
				if currentPrefix.parent.transitions.Empty() {
					currentPrefix.parent.transitions = nil
				}
				currentPrefix = currentPrefix.parent
				trie.internalSize--
			} else {
				var value V
				currentPrefix.value = value
				currentPrefix.flags = currentPrefix.flags & ^(prefixKey | prefixTombstone)
				break
			}
			if (currentPrefix == nil) || currentPrefix.isKey() {
				break
			}
		}
		trie.size--
	}
}

func (trie *trie[K, V]) Get(key []K) (V, error) {
	current := trie.get(key)
	if current != nil {
		return current.value, nil
	}
	var emptyVal V
	return emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Predecessor(key []K) ([]K, V, error) {
	prefix, depth := trie.predecessor(key, false)
	if prefix != nil {
		resultKey, val := makeKey(prefix, depth)
		return resultKey, val, nil
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Successor(key []K) ([]K, V, error) {
	resultKey := make([]K, 0, keyBufferSize)
	next := trie.successor(key, false)
	for current, depth, isKey := next(); current != nil; current, depth, isKey = next() {
		resultKey = resultKey[:depth+1]
		resultKey[depth] = current.parentEdge
		if isKey {
			return resultKey, current.value, nil
		}
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Floor(key []K) ([]K, V, error) {
	prefix, depth := trie.predecessor(key, true)
	if prefix != nil {
		resultKey, val := makeKey(prefix, depth)
		return resultKey, val, nil
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Ceiling(key []K) ([]K, V, error) {
	resultKey := make([]K, 0, keyBufferSize)
	next := trie.successor(key, true)
	for current, depth, isKey := next(); current != nil; current, depth, isKey = next() {
		resultKey = resultKey[:depth+1]
		resultKey[depth] = current.parentEdge
		if isKey {
			return resultKey, current.value, nil
		}
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) First() ([]K, V, error) {
	resultKey := make([]K, 0, keyBufferSize)
	next := trie.forward(trie.root, 0, trie.root)
	for current, depth := next(); current != nil; current, depth = next() {
		resultKey = resultKey[:depth+1]
		resultKey[depth] = current.parentEdge
		if current.isKey() {
			return resultKey, current.value, nil
		}
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Last() ([]K, V, error) {
	if trie.root.transitions != nil && !trie.root.transitions.Empty() {
		prefix, depth := trie.last(trie.root, 0)
		if prefix != nil {
			key, val := makeKey(prefix, depth)
			return key, val, nil
		}
	}
	var emptyVal V
	return nil, emptyVal, errors.New(notFoundMessage)
}

func (trie *trie[K, V]) Size() uint64 {
	return trie.size
}

func (trie *trie[K, V]) InternalSize() uint64 {
	return trie.internalSize
}

// iterators

func (trie *trie[K, V]) All() iter.Seq2[[]K, V] {
	return func(yield func([]K, V) bool) {
		resultKey := make([]K, 0, keyBufferSize)
		next := trie.forward(trie.root, 0, trie.root)
		for current, depth := next(); current != nil; current, depth = next() {
			resultKey = resultKey[:depth+1]
			resultKey[depth] = current.parentEdge
			if current.isKey() && !yield(slices.Clone(resultKey), current.value) {
				return
			}
		}
	}
}

func (trie *trie[K, V]) Prefix(key []K) iter.Seq2[[]K, V] {
	return func(yield func([]K, V) bool) {
		resultKey := make([]K, 0, keyBufferSize)
		next := trie.to(key)
		current, depth := next()
		for ; current != nil && !current.isKey(); current, depth = next() {
			resultKey = resultKey[:depth+1]
			resultKey[depth] = current.parentEdge
			if current.isKey() && !yield(slices.Clone(resultKey), current.value) {
				break
			}
		}
		next = trie.forward(current, depth, current)
		for ; current != nil; current, depth = next() {
			resultKey = resultKey[:depth+1]
			resultKey[depth] = current.parentEdge
			if current.isKey() && !yield(slices.Clone(resultKey), current.value) {
				current = nil
				return
			}
		}
	}
}

func (trie *trie[K, V]) LessThan(key []K, inclusive bool) iter.Seq2[[]K, V] {
	stop, _ := trie.predecessor(key, inclusive)
	return func(yield func([]K, V) bool) {
		resultKey := make([]K, 0, keyBufferSize)
		next := trie.forward(trie.root, 0, trie.root)
		for current, depth := next(); current != nil; current, depth = next() {
			if current == stop {
				return
			}
			resultKey = resultKey[:depth+1]
			resultKey[depth] = current.parentEdge
			if current.isKey() && !yield(slices.Clone(resultKey), current.value) {
				return
			}
		}
	}
}

func (trie *trie[K, V]) GreaterThan(key []K, inclusive bool) iter.Seq2[[]K, V] {
	return func(yield func([]K, V) bool) {
		resultKey := make([]K, 0, keyBufferSize)
		next := trie.successor(key, inclusive)
		for currentPrefix, depth, isKey := next(); currentPrefix != nil; currentPrefix, depth, isKey = next() {
			resultKey = resultKey[:depth+1]
			resultKey[depth] = currentPrefix.parentEdge
			if isKey && !yield(slices.Clone(resultKey), currentPrefix.value) {
				return
			}
		}
	}
}

func (trie *trie[K, V]) Between(loKey []K, loInclusive bool, hiKey []K, hiInclusive bool) iter.Seq2[[]K, V] {
	stop, _ := trie.predecessor(hiKey, hiInclusive)
	return func(yield func([]K, V) bool) {
		resultKey := make([]K, 0, keyBufferSize)
		next := trie.successor(loKey, loInclusive)
		for currentPrefix, depth, isKey := next(); currentPrefix != nil; currentPrefix, depth, isKey = next() {
			if currentPrefix == stop {
				return
			}
			resultKey = resultKey[:depth+1]
			resultKey[depth] = currentPrefix.parentEdge
			if isKey && !yield(slices.Clone(resultKey), currentPrefix.value) {
				return
			}
		}
	}
}
