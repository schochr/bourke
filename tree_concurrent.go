package bourke

import (
	"bytes"
	"cmp"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"iter"

	"github.com/spaolacci/murmur3"
)

// VariableLengthKeyHasher hashes a key. Any type adhering to cmp.Ordered can be passed in. However, the
// encoding is significantly slower than using a fixed length type and hasher.
func VariableLengthKeyHasher[K cmp.Ordered](key K) (uint64, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return 0, err
	}
	hash := murmur3.Sum32(buf.Bytes())
	return uint64(hash), nil
}

// FixedLengthKeyHasher hashes a key. The key must be a fixed length datatype and adhere to cmp.Ordered.
// For example, strings honor the cmp.Ordered constraint but will fail regardless, because they are of
// variable length.
func FixedLengthKeyHasher[K cmp.Ordered](key K) (uint64, error) {
	var buf []byte
	var err error
	if buf, err = binary.Append(buf, binary.BigEndian, key); err != nil {
		return 0, err
	}
	hash := murmur3.Sum32(buf)
	return uint64(hash), nil
}

func (tree *treeConcurrent[K, V]) RLock() {
	for _, shard := range tree.partitions {
		shard.mutex.RLock()
	}
}

func (tree *treeConcurrent[K, V]) RUnlock() {
	for _, shard := range tree.partitions {
		shard.mutex.RUnlock()
	}
}

func (tree *treeConcurrent[K, V]) Put(key K, value V) error {
	shard, err := tree.getShard(key)
	if err != nil {
		return err
	}
	defer shard.mutex.Unlock()
	shard.mutex.Lock()
	_ = shard.tree.Put(key, value)
	return nil
}

func (tree *treeConcurrent[K, V]) Remove(key K) error {
	shard, err := tree.getShard(key)
	if err != nil {
		return err
	}
	defer shard.mutex.Unlock()
	shard.mutex.Lock()
	_ = shard.tree.Remove(key)
	return nil
}

func (tree *treeConcurrent[K, V]) Get(key K) (V, error) {
	shard, err := tree.getShard(key)
	if err != nil {
		var emptyValue V
		return emptyValue, err
	}
	defer shard.mutex.RUnlock()
	shard.mutex.RLock()
	current := shard.tree.get(key)
	if current != nil {
		return current.value, nil
	}
	var emptyValue V
	return emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Successor(key K) (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.successor(key); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.first(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Predecessor(key K) (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.predecessor(key); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.last(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Ceiling(key K) (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.ceiling(key); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.first(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Floor(key K) (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.floor(key); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.last(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyValue V
	return key, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) First() (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.first(); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.first(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Last() (K, V, error) {
	tmp := newTree[K, V]()
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if current := shard.tree.last(); current != nil {
			_ = tmp.Put(current.key, current.value)
		}
	}
	if smallest := tmp.last(); smallest != nil {
		return smallest.key, smallest.value, nil
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, errors.New("not_found")
}

func (tree *treeConcurrent[K, V]) Empty() bool {
	tree.RLock()
	defer tree.RUnlock()
	for _, shard := range tree.partitions {
		if !shard.tree.Empty() {
			return false
		}
	}
	return true
}

// iterators

func (tree *treeConcurrent[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		tmp := newTree[K, func() (K, V, bool)]()
		for _, shard := range tree.partitions {
			next, stop := iter.Pull2(shard.tree.All())
			key, _, valid := next()
			if valid {
				next, stop = iter.Pull2(shard.tree.All())
				_ = tmp.Put(key, next)
			}
			defer stop()
		}
		current := tmp.first()
		for current != nil {
			_ = tmp.Remove(current.key)
			key, value, valid := current.value()
			if !valid {
				current = tmp.first()
			} else if yield(key, value) {
				_ = tmp.Put(key, current.value)
				current = tmp.first()
			}
		}
	}
}

func (tree *treeConcurrent[K, V]) LessThan(hiKey K, inclusive bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		tmp := newTree[K, func() (K, V, bool)]()
		for _, shard := range tree.partitions {
			next, stop := iter.Pull2(shard.tree.LessThan(hiKey, inclusive))
			key, _, valid := next()
			if valid {
				next, stop = iter.Pull2(shard.tree.LessThan(hiKey, inclusive))
				_ = tmp.Put(key, next)
			}
			defer stop()
		}
		current := tmp.first()
		for current != nil {
			_ = tmp.Remove(current.key)
			key, value, valid := current.value()
			if !valid {
				current = tmp.first()
			} else if yield(key, value) {
				_ = tmp.Put(key, current.value)
				current = tmp.first()
			}
		}
	}
}

func (tree *treeConcurrent[K, V]) GreaterThan(loKey K, inclusive bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		tmp := newTree[K, func() (K, V, bool)]()
		for _, shard := range tree.partitions {
			next, stop := iter.Pull2(shard.tree.GreaterThan(loKey, inclusive))
			key, _, valid := next()
			if valid {
				next, stop = iter.Pull2(shard.tree.GreaterThan(loKey, inclusive))
				_ = tmp.Put(key, next)
			}
			defer stop()
		}
		current := tmp.first()
		for current != nil {
			_ = tmp.Remove(current.key)
			key, value, valid := current.value()
			if !valid {
				current = tmp.first()
			} else if yield(key, value) {
				_ = tmp.Put(key, current.value)
				current = tmp.first()
			}
		}
	}
}

func (tree *treeConcurrent[K, V]) Between(loKey K, inclusiveLo bool, hiKey K, inclusiveHi bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		tmp := newTree[K, func() (K, V, bool)]()
		for _, shard := range tree.partitions {
			next, stop := iter.Pull2(shard.tree.Between(loKey, inclusiveLo, hiKey, inclusiveHi))
			key, _, valid := next()
			if valid {
				next, stop = iter.Pull2(shard.tree.Between(loKey, inclusiveLo, hiKey, inclusiveHi))
				_ = tmp.Put(key, next)
			}
			defer stop()
		}
		current := tmp.first()
		for current != nil {
			_ = tmp.Remove(current.key)
			key, value, valid := current.value()
			if !valid {
				current = tmp.first()
			} else if yield(key, value) {
				_ = tmp.Put(key, current.value)
				current = tmp.first()
			}
		}
	}
}
