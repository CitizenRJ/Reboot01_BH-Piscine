package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	return searchItemHelper(root, elem, nil)
}

func searchItemHelper(root *TreeNode, elem string, parent *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return parent
	}
	if elem < root.Data {
		return searchItemHelper(root.Left, elem, root)
	} else {
		return searchItemHelper(root.Right, elem, root)
	}
}
