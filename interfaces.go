package bourke

import (
	"cmp"
	"iter"
)

// A Stack is a LIFO data structure with O(1) runtime complexity for all operations.
type Stack[V any] interface {
	// Push adds a new element to the top.
	Push(value V) error
	// Pop returns and removes the top element.
	Pop() (V, error)
	// Peek returns the top element without removing it.
	Peek() (V, error)
	// Size returns the number of elements currently on the stack.
	Size() int
	// Reset clears the stack (size 0).
	Reset()
}

// Tree implements the Red-Black tree data structure, an ordered map with O(log n) runtime complexity for common operations. Invariants:
// "The Red-Black invariant is that every path from the root to a leaf has the same number of black nodes, and no such path has two red
// nodes parent a row. Thus, each leaf is at most twice as deep as any other leaf, and this means that the height of an N-node tree is
// at most 2 log N" [Efficient Verified Red-Black Trees | Andrew W. Appel | Princeton University, Princeton NJ 08540, USA
// https://www.cs.princeton.edu/~appel/papers/redblack.pdf]
//
// https://en.wikipedia.org/wiki/Red%E2%80%93black_tree
// https://ccs.neu.edu/~camoy/pub/red-black-tree.pdf
type Tree[K cmp.Ordered, V any] interface {
	// Put adds a new key-value pair.
	Put(key K, value V) error
	// Remove (hard) deletes the key-value pair associated with the input key.
	Remove(key K) error

	// Get returns the value associated with the input key.
	Get(key K) (V, error)
	// Predecessor returns the key-value pair associated with the biggest key that is strictly less than the input key.
	Predecessor(key K) (K, V, error)
	// Successor returns the key-value pair associated with the smallest key that is strictly greater than the input key.
	Successor(key K) (K, V, error)
	// Floor returns the key-value pair with the biggest key that is equal (takes precedence) or less than the input key.
	Floor(key K) (K, V, error)
	// Ceiling returns the key-value pair with the smallest key that is equal (takes precedence) or greater than the input key.
	Ceiling(key K) (K, V, error)
	// First returns the key-value pair with the smallest key.
	First() (K, V, error)
	// Last returns the key-value pair with the biggest key.
	Last() (K, V, error)

	// All returns an in-key-order iterator over they entire map.
	All() iter.Seq2[K, V]
	// LessThan returns an in-key-order iterator over key-value pairs that are (strictly) less than the input key.
	LessThan(key K, inclusive bool) iter.Seq2[K, V]
	// GreaterThan returns an in-key-order iterator over key-value pairs that are (strictly) greater than the input key.
	GreaterThan(key K, inclusive bool) iter.Seq2[K, V]
	// Between returns an in-key-order iterator over key-value pairs that are (strictly) greater
	// than the lower bound key AND (strictly) less than the upper bound key.
	Between(loKey K, loInclusive bool, hiKey K, hiInclusive bool) iter.Seq2[K, V]

	// Empty returns true if the root is nil, implying a size of 0 (zero), false otherwise.
	Empty() bool
}

// Trie implements a Prefix Tree. 'K' does not refer to the type of the key itself but rather its individual parts that must adhere to
// cmp.Ordered. For example, a string's UTF-8 byte representation, i.e. `[]byte("string")` (the default) can be stored with the expected
// lexicographical ordering.
//
// https://en.wikipedia.org/wiki/Trie
// https://www.vldb.org/pvldb/vol15/p3359-lambov.pdf
// https://github.com/apache/cassandra/blob/76d0c25139fb1acfae616e3c750887238378d467/src/java/org/apache/cassandra/utils/bytecomparable/ByteComparable.md
// https://github.com/apache/cassandra/blob/76d0c25139fb1acfae616e3c750887238378d467/src/java/org/apache/cassandra/db/tries/Trie.md
// https://github.com/apache/cassandra/blob/76d0c25139fb1acfae616e3c750887238378d467/src/java/org/apache/cassandra/db/tries/InMemoryTrie.md
type Trie[K cmp.Ordered, V any] interface {
	// Put adds a new key-value pair.
	Put(key []K, value V)
	// Remove (hard) deletes the key-value pair associated with the input key.
	Remove(key []K)
	// Tombstone will always create the key for the tombstone marker even if an "active" key for it does not exist. They will not affect the
	// size of the trie. However, its key parts will be reflected in the internal size of the trie. It's a form of a soft delete that doesn't
	// require the existence of the associated key.
	Tombstone(key []K)

	// Get returns the value associated with the input key.
	Get(key []K) (V, error)
	// Predecessor returns the key-value pair associated with the greatest key that is strictly less than the input key.
	Predecessor(key []K) ([]K, V, error)
	// Successor returns the key-value pair associated with the smallest key that is strictly greater than the input key.
	Successor(key []K) ([]K, V, error)
	// Floor returns the key-value pair with the greatest key that is equal (takes precedence) or less than the input key.
	Floor(key []K) ([]K, V, error)
	// Ceiling returns the key-value pair with the smallest key that is equal (takes precedence) or greater than the input key.
	Ceiling(key []K) ([]K, V, error)
	// First returns the key-value pair with the smallest key.
	First() ([]K, V, error)
	// Last returns the key-value pair with the greatest key.
	Last() ([]K, V, error)

	// All returns an in-key-order iterator over they entire map.
	All() iter.Seq2[[]K, V]
	// LessThan returns an in-key-order iterator over key-value pairs that are (strictly) less than the input key.
	LessThan(key []K, inclusive bool) iter.Seq2[[]K, V]
	// GreaterThan returns an in-key-order iterator over key-value pairs that are (strictly) greater than the input key.
	GreaterThan(key []K, inclusive bool) iter.Seq2[[]K, V]
	// Between returns an in-key-order iterator over key-value pairs that are (strictly) greater
	// than the lower bound key AND (strictly) less than the upper bound key.
	Between(loKey []K, loInclusive bool, hiKey []K, hiInclusive bool) iter.Seq2[[]K, V]
	// Prefix returns an in-key-order iterator over key-value pairs that contain the input key as prefix.
	Prefix(key []K) iter.Seq2[[]K, V]

	// Size returns the number of keys (excluding tombstone marker) in the Trie.
	Size() uint64
	// InternalSize refers to the actual number of vertices (key parts) in the Trie including tombstone markers.
	InternalSize() uint64
}

// ReadSync provides RW Locking over ALL partitions.
type ReadSync interface {
	// RLock read-locks all partitions. This is necessary if (read) synchronization is required for iterators, which can't lock because it is
	// not known whether the caller blocks or otherwise introduces latency between iterations.
	RLock()
	// RUnlock read-unlocks all partitions. See RLock.
	RUnlock()
}

// NewStack creates a new, empty stack instance. If size is known in advance an initial capacity
// should be set to avoid re-sizing of the internal slice.
func NewStack[T any](initialCapacity int) Stack[T] {
	return &stack[T]{index: -1, array: make([]T, initialCapacity)}
}

// NewTree creates a new, empty Red-Black Tree instance.
func NewTree[K cmp.Ordered, V any]() Tree[K, V] {
	return newTree[K, V]()
}

// NewTreeConcurrent creates a thread-safe Red-Black Tree instance with number of partitions (partition factor).
// In highly concurrent scenarios more partitions can improve performance, particularly for writes. However, the
// trade-off is slower reads.
func NewTreeConcurrent[K cmp.Ordered, V any](partitions uint32, hasher func(key K) (uint64, error)) interface {
	Tree[K, V]
	ReadSync
} {
	return newTreeConcurrent[K, V](partitions, hasher)
}

// NewTrie creates a new, empty Trie instance.
func NewTrie[K cmp.Ordered, V any]() Trie[K, V] {
	return newTrie[K, V]()
}
