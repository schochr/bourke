package bourke

import (
	"cmp"
	"sync"
)

type shard[K cmp.Ordered, V any] struct {
	tree  *tree[K, V]
	mutex *sync.RWMutex
}

type treeConcurrent[K cmp.Ordered, V any] struct {
	partitions []*shard[K, V]
	hasher     func(key K) (uint64, error)
}

func newTreeConcurrent[K cmp.Ordered, V any](partitions uint32, hasher func(key K) (uint64, error)) *treeConcurrent[K, V] {
	shards := make([]*shard[K, V], partitions)
	var i uint32 = 0
	for ; i < partitions; i++ {
		shards[i] = &shard[K, V]{&tree[K, V]{root: nil}, new(sync.RWMutex)}
	}
	return &treeConcurrent[K, V]{shards, hasher}
}

func (tree *treeConcurrent[K, V]) getShard(key K) (*shard[K, V], error) {
	hash, err := tree.hasher(key)
	if err != nil {
		return nil, err
	}
	return tree.partitions[hash%uint64(len(tree.partitions))], nil
}
