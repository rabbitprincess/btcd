package treap

import "sync"

var treapNodePool = sync.Pool{
	New: func() interface{} {
		return &treapNode{}
	},
}

func newTreapNode(key, value []byte, priority int) *treapNode {
	node := treapNodePool.Get().(*treapNode)
	node.key = key
	node.value = value
	node.priority = priority
	node.left = nil
	node.right = nil
	return node
}

func releaseTreapNode(node *treapNode) {
	if node.key != nil {
		node.key = node.key[:0]
	}
	if node.value != nil {
		node.value = node.value[:0]
	}
	node.priority = 0
	node.left = nil
	node.right = nil
	treapNodePool.Put(node)
}

func cloneTreapNode(node *treapNode) *treapNode {
	// Reuse a node from the pool or create a new one
	retNode := treapNodePool.Get().(*treapNode)

	// Copy fields
	retNode.key = node.key
	retNode.value = node.value
	retNode.priority = node.priority
	retNode.left = node.left
	retNode.right = node.right

	return retNode
}
