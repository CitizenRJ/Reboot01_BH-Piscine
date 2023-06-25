package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	if node.Left != nil {
		node.Left.Parent = rplc
	}
	if node.Right != nil {
		node.Right.Parent = rplc
	}
	return root
}
