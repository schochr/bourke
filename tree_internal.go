package bourke

import (
	"cmp"
)

type tree[K cmp.Ordered, V any] struct {
	root *vertex[K, V]
}

type vertex[K cmp.Ordered, V any] struct {
	isBlack bool
	parent  *vertex[K, V]
	lt      *vertex[K, V]
	gt      *vertex[K, V]
	key     K
	value   V
}

func newTree[K cmp.Ordered, V any]() *tree[K, V] {
	return &tree[K, V]{root: nil}
}

func (tree *tree[K, V]) get(key K) *vertex[K, V] {
	for current := tree.root; current != nil; {
		if key < current.key {
			current = current.lt
		} else if key > current.key {
			current = current.gt
		} else {
			return current
		}
	}
	return nil
}

func (tree *tree[K, V]) successor(key K) *vertex[K, V] {
	for current := tree.root; current != nil; {
		if key < current.key {
			if current.lt == nil {
				return current
			}
			current = current.lt
		} else {
			if current.gt == nil {
				parentVertex := current.parent
				for parentVertex != nil && current == parentVertex.gt {
					current = parentVertex
					parentVertex = parentVertex.parent
				}
				return parentVertex
			}
			current = current.gt
		}
	}
	return nil
}

func (tree *tree[K, V]) predecessor(key K) *vertex[K, V] {
	current := tree.root
	for current != nil {
		if key < current.key {
			current = current.lt
		} else if key > current.key {
			if current.gt == nil {
				return current
			}
			if key < current.gt.key {
				break
			}
			current = current.gt
		} else {
			if current.lt != nil {
				current = current.lt
				for current.gt != nil {
					current = current.gt
				}
				return current
			}
			parentVertex := current.parent
			for parentVertex != nil && current == parentVertex.lt {
				current = parentVertex
				parentVertex = parentVertex.parent
			}
			return parentVertex
		}
	}
	return current
}

func (tree *tree[K, V]) ceiling(key K) *vertex[K, V] {
	for current := tree.root; current != nil; {
		if key < current.key {
			if current.lt == nil {
				return current
			}
			current = current.lt
		} else if key == current.key {
			return current
		} else {
			if current.gt == nil {
				parentVertex := current.parent
				for parentVertex != nil && current == parentVertex.gt {
					current = parentVertex
					parentVertex = parentVertex.parent
				}
				return parentVertex
			}
			current = current.gt
		}
	}
	return nil
}

func (tree *tree[K, V]) floor(key K) *vertex[K, V] {
	for current := tree.root; current != nil; {
		if key > current.key {
			if current.gt == nil {
				return current
			}
			current = current.gt
		} else if key == current.key {
			return current
		} else {
			if current.lt == nil {
				parentVertex := current.parent
				for parentVertex != nil && current == parentVertex.lt {
					current = parentVertex
					parentVertex = parentVertex.parent
				}
				return parentVertex
			}
			current = current.lt
		}
	}
	return nil
}

func (tree *tree[K, V]) first() *vertex[K, V] {
	current := tree.root
	for current != nil && current.lt != nil {
		current = current.lt
	}
	return current
}

func (tree *tree[K, V]) last() *vertex[K, V] {
	current := tree.root
	for current != nil && current.gt != nil {
		current = current.gt
	}
	return current
}

func (current *vertex[K, V]) from() func() *vertex[K, V] {
	return func() *vertex[K, V] {
		if current != nil {
			if current.gt != nil {
				current = current.gt
				for current.lt != nil {
					current = current.lt
				}
			} else {
				parentVertex := current.parent
				for parentVertex != nil && current == parentVertex.gt {
					current = parentVertex
					parentVertex = parentVertex.parent
				}
				current = parentVertex
			}
		}
		return current
	}
}
