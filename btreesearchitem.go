package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	leftResult := BTreeSearchItem(root.Left, elem)
	if leftResult != nil {
		return leftResult
	}
	rightResult := BTreeSearchItem(root.Right, elem)
	return rightResult
}
