//go:build !solution

package treeiter

type TreeNode[T any] interface {
	Left() *T
	Right() *T
}

func DoInOrder[N TreeNode[N]](node *N, visit func(*N)) {
	if node == nil {
		return
	}
	DoInOrder((*node).Left(), visit)
	visit(node)
	DoInOrder((*node).Right(), visit)
}
