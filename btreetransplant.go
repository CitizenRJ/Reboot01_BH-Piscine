package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	new := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		new.Parent.Left = rplc
	} else {
		new.Parent.Right = rplc
	}
	new.Parent = node.Parent
	return root
}
