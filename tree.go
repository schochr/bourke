package bourke

import (
	"errors"
	"iter"
)

func (tree *tree[K, V]) Put(key K, value V) error {
	if tree.root == nil {
		tree.root = &vertex[K, V]{isBlack: true, parent: nil, lt: nil, gt: nil, key: key, value: value}
		return nil
	}
	current := tree.root
	for current != nil {
		if key < current.key {
			if current.lt == nil {
				current.lt = &vertex[K, V]{isBlack: false, parent: current, lt: nil, gt: nil, key: key, value: value}
				tree.balancePut(current.lt)
				return nil
			}
			current = current.lt
		} else if key > current.key {
			if current.gt == nil {
				current.gt = &vertex[K, V]{isBlack: false, parent: current, lt: nil, gt: nil, key: key, value: value}
				tree.balancePut(current.gt)
				return nil
			}
			current = current.gt
		} else {
			current.value = value
			return nil
		}
	}
	return nil
}

func (tree *tree[K, V]) Remove(key K) error {
	current := tree.get(key)
	if current == nil {
		return nil
	}
	// two children (pre-processing)
	if (current.lt != nil) && (current.gt != nil) {
		tmp := current.from()() // next will not produce <nil>, due to current.gt != nil being checked before invocation
		current.key = tmp.key
		current.value = tmp.value
		current = tmp
	}
	movingVertex := current.gt
	if current.lt != nil {
		movingVertex = current.lt
	}
	if movingVertex != nil { // has at least one child
		movingVertex.parent = current.parent
		if current.parent == nil {
			tree.root = movingVertex
		} else if current == current.parent.lt {
			current.parent.lt = movingVertex
		} else {
			current.parent.gt = movingVertex
		}
		current.parent = nil
		current.lt = nil
		current.gt = nil
		if current.isBlack == true {
			tree.balanceRemove(movingVertex)
		}
	} else if current.parent == nil { // No children AND (current == tree.root) implies (tree.size == 1), hence unsetting root is sufficient.
		tree.root = nil
	} else { // no children
		if current.isBlack == true {
			tree.balanceRemove(current)
		}
		if current.parent != nil {
			if current == current.parent.lt {
				current.parent.lt = nil
			} else if current == current.parent.gt {
				current.parent.gt = nil
			}
			current.parent = nil
		}
	}
	return nil
}

func (tree *tree[K, V]) Get(key K) (V, error) {
	targetVertex := tree.get(key)
	if targetVertex != nil {
		return targetVertex.value, nil
	}
	var emptyValue V
	return emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Successor(key K) (K, V, error) {
	current := tree.successor(key)
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Predecessor(key K) (K, V, error) {
	current := tree.predecessor(key)
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Ceiling(key K) (K, V, error) {
	current := tree.ceiling(key)
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Floor(key K) (K, V, error) {
	current := tree.floor(key)
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) First() (K, V, error) {
	current := tree.first()
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Last() (K, V, error) {
	current := tree.last()
	if current != nil {
		return current.key, current.value, nil
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, errors.New("not_found")
}

func (tree *tree[K, V]) Empty() bool {
	return tree.root == nil
}

// iterators

func (tree *tree[K, V]) All() iter.Seq2[K, V] {
	current := tree.first()
	next := current.from()
	return func(yield func(K, V) bool) {
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = next()
		}
	}
}

func (tree *tree[K, V]) LessThan(hiKey K, inclusive bool) iter.Seq2[K, V] {
	current := tree.first()
	next := current.from()
	return func(yield func(K, V) bool) {
		for current != nil {
			if (!inclusive && current.key >= hiKey) || (inclusive && current.key > hiKey) || !yield(current.key, current.value) {
				return
			}
			current = next()
		}
	}
}

func (tree *tree[K, V]) GreaterThan(loKey K, inclusive bool) iter.Seq2[K, V] {
	var current *vertex[K, V]
	if inclusive {
		current = tree.ceiling(loKey)
	} else {
		current = tree.successor(loKey)
	}
	next := current.from()
	return func(yield func(K, V) bool) {
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = next()
		}
	}
}

func (tree *tree[K, V]) Between(loKey K, inclusiveLo bool, hiKey K, inclusiveHi bool) iter.Seq2[K, V] {
	var current *vertex[K, V]
	if inclusiveLo {
		current = tree.ceiling(loKey)
	} else {
		current = tree.successor(loKey)
	}
	next := current.from()
	return func(yield func(K, V) bool) {
		for current != nil {
			if (!inclusiveHi && current.key >= hiKey) || (inclusiveHi && current.key > hiKey) || !yield(current.key, current.value) {
				return
			}
			current = next()
		}
	}
}
