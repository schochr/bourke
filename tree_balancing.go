package bourke

func (tree *tree[K, V]) rotateLeft(vertex *vertex[K, V]) {
	tmp := vertex.gt
	vertex.gt = tmp.lt
	if tmp.lt != nil {
		tmp.lt.parent = vertex
	}
	tmp.parent = vertex.parent
	if vertex.parent == nil {
		tree.root = tmp
	} else if vertex.parent.lt == vertex {
		vertex.parent.lt = tmp
	} else {
		vertex.parent.gt = tmp
	}
	tmp.lt = vertex
	vertex.parent = tmp
}

func (tree *tree[K, V]) rotateRight(vertex *vertex[K, V]) {
	tmp := vertex.lt
	vertex.lt = tmp.gt
	if tmp.gt != nil {
		tmp.gt.parent = vertex
	}
	tmp.parent = vertex.parent
	if vertex.parent == nil {
		tree.root = tmp
	} else if vertex.parent.gt == vertex {
		vertex.parent.gt = tmp
	} else {
		vertex.parent.lt = tmp
	}
	tmp.gt = vertex
	vertex.parent = tmp
}

func (tree *tree[K, V]) balancePut(vertex *vertex[K, V]) {
	for (vertex.parent != nil) && (vertex.parent.parent != nil) && (vertex.parent.isBlack == false) {
		if vertex.parent == vertex.parent.parent.lt {
			gtUncle := vertex.parent.parent.gt
			if (gtUncle != nil) && (gtUncle.isBlack == false) {
				gtUncle.isBlack = true
				vertex.parent.isBlack = true
				vertex.parent.parent.isBlack = false
				vertex = vertex.parent.parent
			} else {
				if vertex == vertex.parent.gt {
					vertex = vertex.parent
					tree.rotateLeft(vertex)
				}
				vertex.parent.isBlack = true
				vertex.parent.parent.isBlack = false
				if vertex.parent.parent != nil {
					tree.rotateRight(vertex.parent.parent)
				}
			}
		} else {
			ltUncle := vertex.parent.parent.lt
			if (ltUncle != nil) && (ltUncle.isBlack == false) {
				ltUncle.isBlack = true
				vertex.parent.isBlack = true
				vertex.parent.parent.isBlack = false
				vertex = vertex.parent.parent
			} else {
				if vertex == vertex.parent.lt {
					vertex = vertex.parent
					tree.rotateRight(vertex)
				}
				vertex.parent.isBlack = true
				vertex.parent.parent.isBlack = false
				if vertex.parent.parent != nil {
					tree.rotateLeft(vertex.parent.parent)
				}
			}
		}
	}
	tree.root.isBlack = true
}

func (tree *tree[K, V]) balanceRemove(vertex *vertex[K, V]) {
	for (vertex != tree.root) && (vertex.isBlack == true) {
		if vertex == vertex.parent.lt {
			sibling := vertex.parent.gt
			if (sibling != nil) && (sibling.isBlack == false) {
				sibling.isBlack = true
				vertex.parent.isBlack = false
				tree.rotateLeft(vertex.parent)
				sibling = vertex.parent.gt
			}
			if (sibling == nil) ||
				(((sibling.lt == nil) || (sibling.lt.isBlack == true)) &&
					((sibling.gt == nil) || (sibling.gt.isBlack == true))) {
				if sibling != nil {
					sibling.isBlack = false
				}
				vertex = vertex.parent // continue loop
			} else {
				if (sibling.gt == nil) || (sibling.gt.isBlack == true) {
					if sibling.lt != nil {
						sibling.lt.isBlack = true
					}
					sibling.isBlack = false
					tree.rotateRight(sibling)
					sibling = vertex.parent.gt
				}
				sibling.isBlack = vertex.parent.isBlack
				vertex.parent.isBlack = true
				sibling.gt.isBlack = true
				tree.rotateLeft(vertex.parent)
				break
			}
		} else {
			sibling := vertex.parent.lt
			if (sibling != nil) && (sibling.isBlack == false) {
				sibling.isBlack = true
				vertex.parent.isBlack = false
				tree.rotateRight(vertex.parent)
				sibling = vertex.parent.lt
			}
			if (sibling == nil) ||
				(((sibling.lt == nil) || (sibling.lt.isBlack == true)) &&
					((sibling.gt == nil) || (sibling.gt.isBlack == true))) {
				if sibling != nil {
					sibling.isBlack = false
				}
				vertex = vertex.parent // continue loop
			} else {
				if (sibling.lt == nil) || (sibling.lt.isBlack == true) {
					if sibling.gt != nil {
						sibling.gt.isBlack = true
					}
					sibling.isBlack = false
					tree.rotateLeft(sibling)
					sibling = vertex.parent.lt
				}
				sibling.isBlack = vertex.parent.isBlack
				vertex.parent.isBlack = true
				sibling.lt.isBlack = true
				tree.rotateRight(vertex.parent)
				break
			}
		}
	}
	vertex.isBlack = true
}
