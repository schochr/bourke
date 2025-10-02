package bourke

import (
	"cmp"
)

type vertexInfo[K cmp.Ordered, V any] struct {
	vertex      *vertex[K, V]
	totalHeight int
	blackHeight int
}

func computeHeights[K cmp.Ordered, V any](vertex *vertex[K, V]) vertexInfo[K, V] {
	totalHeight := 0
	blackHeight := 0
	currentVertex := vertex
	tmp := vertex
	for vertex != nil {
		if totalHeight > 0 {
			if (currentVertex.isBlack == false) && (vertex.isBlack == false) {
				panic("invariant violation: 'no adjacent red vertices'")
			} else {
				currentVertex = vertex
			}
		}
		totalHeight++
		if vertex.isBlack == true {
			blackHeight++
		}
		vertex = vertex.parent
	}
	return vertexInfo[K, V]{totalHeight: totalHeight, blackHeight: blackHeight, vertex: tmp}
}

func calculateHeights[K cmp.Ordered, V any](tree *tree[K, V]) map[K]vertexInfo[K, V] {
	heightMap := make(map[K]vertexInfo[K, V])
	traversal := NewStack[*vertex[K, V]](32)
	traversal.Push(tree.root)
	for traversal.Size() > 0 {
		current, _ := traversal.Pop()
		heightMap[current.key] = computeHeights(current)
		if (*current).lt != nil {
			traversal.Push((*current).lt)
		}
		if (*current).gt != nil {
			traversal.Push((*current).gt)
		}
	}
	return heightMap
}

func verifyInvariants[K cmp.Ordered, V any](tree *tree[K, V]) map[K]vertexInfo[K, V] {
	if tree.root.isBlack == false {
		panic("invariant violation: 'root is black'")
	}
	vertexInfoMap := calculateHeights(tree)
	currentLeafBlackHeight := -1
	for _, info := range vertexInfoMap {
		if (info.vertex.lt == nil) && (info.vertex.gt == nil) {
			if (currentLeafBlackHeight != info.blackHeight) && (currentLeafBlackHeight > -1) {
				panic("invariant violation: 'equal leaf black-height'")
			}
			currentLeafBlackHeight = info.blackHeight
		}
	}
	return vertexInfoMap
}
