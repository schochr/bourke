package bourke

import (
	"cmp"
)

const (
	prefixEmpty     byte = 0b_0000_0000
	prefixKey       byte = 0b_0000_0001
	prefixTombstone byte = 0b_0000_0010
)

const keyBufferSize byte = 64

const notFoundMessage string = "NOT_FOUND"

type trie[K cmp.Ordered, V any] struct {
	root         *prefix[K, V]
	size         uint64 // Logical size of the trie. Does not account for tombstones.
	internalSize uint64 // Number of actual vertices in the tree including those leading to tombstones with no further transitions.
}

// A prefix represents a vertex in the trie. Vertices in a trie do not directly correspond to a key as their parts reside
// on the edges i.e. the transitions. A prefix should be interpreted as the path (set of vertices) taken to get to
// said vertex.
type prefix[K cmp.Ordered, V any] struct {
	parent      *prefix[K, V]
	parentEdge  K
	transitions *tree[K, *prefix[K, V]]
	flags       byte
	value       V
}

func newTrie[K cmp.Ordered, V any]() *trie[K, V] {
	var empty K
	return &trie[K, V]{
		root: &prefix[K, V]{
			parent:      nil,
			parentEdge:  empty,
			transitions: nil,
			flags:       prefixEmpty,
		},
		size:         0,
		internalSize: 0,
	}
}

func (prefix *prefix[K, V]) isKey() bool {
	return prefix.flags&prefixKey == prefixKey
}

func (prefix *prefix[K, V]) isTombstone() bool {
	return prefix.flags&prefixTombstone == prefixTombstone
}

func (trie *trie[K, V]) makeInternalPrefix(currentPrefix *prefix[K, V], keyByte K) *prefix[K, V] {
	if currentPrefix.transitions == nil {
		transitions := newTree[K, *prefix[K, V]]()
		newPrefix := &prefix[K, V]{parent: currentPrefix, parentEdge: keyByte, transitions: nil, flags: prefixEmpty}
		_ = transitions.Put(keyByte, newPrefix)
		currentPrefix.transitions = transitions
		trie.internalSize++
		return newPrefix
	}
	transition := currentPrefix.transitions.get(keyByte)
	if transition == nil {
		newPrefix := &prefix[K, V]{parent: currentPrefix, parentEdge: keyByte, transitions: nil, flags: prefixEmpty}
		_ = currentPrefix.transitions.Put(keyByte, newPrefix)
		trie.internalSize++
		return newPrefix
	}
	return transition.value
}

func (trie *trie[K, V]) get(key []K) *prefix[K, V] {
	resultKey := make([]K, 0, keyBufferSize)
	next := trie.to(key)
	lastIndex := uint(len(key) - 1)
	for current, depth := next(); current != nil; current, depth = next() {
		resultKey = resultKey[:depth+1]
		resultKey[depth] = current.parentEdge
		if current.isKey() && depth == lastIndex {
			return current
		}
	}
	return nil
}

func (trie *trie[K, V]) successor(key []K, inclusive bool) func() (*prefix[K, V], uint, bool) {
	var depth uint = 0
	current := trie.root
	var next func() (*prefix[K, V], uint) = nil
	return func() (*prefix[K, V], uint, bool) {
		if next != nil { // this is set once the input key is exhausted or not matching at any point
			current, depth := next()
			if current != nil {
				return current, depth, current.isKey()
			}
			return nil, 0, false // this only happens when iterating past the calculated/external `isKey`
		}
		if current == nil {
			return nil, 0, false
		}
		var transition *vertex[K, *prefix[K, V]] = nil
		if current.transitions != nil {
			if len(key) == int(depth) {
				next = trie.forward(current, depth, trie.root)
				current, depth = next()
				return current, depth, current.isKey()
			}
			transition = current.transitions.ceiling(key[depth])
		}
		if transition != nil {
			current = transition.value
			if transition.key != key[depth] {
				next = trie.forward(transition.value, depth, trie.root)
				return transition.value, depth, transition.value.isKey()
			}
			depth++
			if inclusive && len(key) == int(depth) {
				return transition.value, depth - 1, transition.value.isKey()
			}
			return transition.value, depth - 1, false
		} else if current.parent != nil {
			depth--
			for {
				transition := current.parent.transitions.successor(current.parentEdge)
				if transition != nil {
					next = trie.forward(transition.value, depth, trie.root)
					return transition.value, depth, transition.value.isKey()
				}
				current = current.parent
				depth--
				if current.parent == nil {
					break
				}
			}
		}
		return nil, 0, false
	}
}

func (trie *trie[K, V]) last(current *prefix[K, V], depth uint) (*prefix[K, V], uint) {
	for current != nil {
		var transition *vertex[K, *prefix[K, V]] = nil
		if current.transitions != nil {
			transition = current.transitions.last()
		}
		if transition != nil {
			if current != trie.root {
				depth++
			}
			if transition.value.transitions == nil && transition.value.isKey() {
				return transition.value, depth
			}
			current = transition.value
		} else if current.parent != nil {
			for current.parent != nil {
				if current.isKey() {
					return current, depth
				}
				if transition = current.parent.transitions.predecessor(current.parentEdge); transition != nil {
					current = transition.value
					break
				}
				if current.parent == nil {
					break
				}
				depth--
				current = current.parent
			}
		}
		transition = nil
	}
	return current, depth
}

func (trie *trie[K, V]) keyPredecessor(key []K) (*prefix[K, V], uint, bool) {
	var depth uint = 0
	current := trie.root
	lastIndex := uint(len(key)) - 1
	for current != nil && current.transitions != nil {
		if transition := current.transitions.floor(key[depth]); transition != nil {
			if transition.key != key[depth] {
				return transition.value, depth, false
			}
			if depth == lastIndex {
				return transition.value, depth, true
			}
			depth++
			current = transition.value
		} else {
			if depth == lastIndex && current == trie.root {
				return nil, 0, false
			}
			break
		}
	}
	return current, depth - 1, true
}

func makeKey[K cmp.Ordered, V any](current *prefix[K, V], depth uint) ([]K, V) {
	key := make([]K, depth+1)
	val := current.value
	for current != nil && current.parent != nil {
		key[depth] = current.parentEdge
		depth--
		current = current.parent
	}
	return key, val
}

func (trie *trie[K, V]) predecessor(key []K, inclusive bool) (*prefix[K, V], uint) {
	prefix, depth, matches := trie.keyPredecessor(key)
	if prefix == nil {
		return nil, 0
	}
	if matches {
		if prefix.isKey() {
			if uint(len(key))-1 == depth {
				if inclusive {
					return prefix, depth
				}
			} else if uint(len(key))-1 > depth {
				return prefix, depth
			}
		}
		for prefix != nil && prefix.parent != nil {
			if prefix.parent.isKey() {
				return prefix.parent, depth - 1
			}
			pred := prefix.parent.transitions.predecessor(prefix.parentEdge)
			if pred != nil {
				prefix, depth = trie.last(pred.value, depth)
				return prefix, depth
			}
			depth--
			prefix = prefix.parent
		}
		return nil, 0
	}
	prefix, depth = trie.last(prefix, depth)
	return prefix, depth
}

func (trie *trie[K, V]) to(key []K) func() (*prefix[K, V], uint) {
	if key == nil || len(key) == 0 {
		return func() (*prefix[K, V], uint) { return nil, 0 }
	}
	var depth uint = 0
	current := trie.root
	lastIndex := uint(len(key) - 1)
	return func() (*prefix[K, V], uint) {
		if current != nil && current.transitions != nil && depth <= lastIndex {
			transition := current.transitions.get(key[depth])
			if transition != nil {
				if depth == lastIndex {
					current = nil
				} else {
					current = transition.value
				}
				tmpDepth := depth
				depth++
				return transition.value, tmpDepth // 0 (zero) actually corresponds to the edge between root and the transition found.
			}
		}
		current = nil
		return nil, 0
	}
}

// `depth` MUST correspond to `current`. `depth` is not a stored property and must be provided.
// It is not needed for traversal, however it is essential to reconstruct keys.
func (trie *trie[K, V]) forward(
	current *prefix[K, V],
	depth uint,
	backstop *prefix[K, V],
) func() (*prefix[K, V], uint) {
	return func() (*prefix[K, V], uint) {
		if current != nil {
			var transition *vertex[K, *prefix[K, V]] = nil
			if current.transitions != nil {
				transition = current.transitions.first()
			}
			if transition != nil {
				if current != trie.root {
					depth++
				}
				current = transition.value
			} else if current != backstop {
				for {
					if transition = current.parent.transitions.successor(current.parentEdge); transition != nil {
						current = transition.value
						break
					}
					depth--
					if current = current.parent; current == backstop {
						break
					}
				}
			}
			transition = nil
		}
		if current != backstop {
			return current, depth
		}
		return nil, 0
	}
}
