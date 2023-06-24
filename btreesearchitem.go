package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.data == elem {
		return root
	} else if root.data > elem {
		return BTreeSearchItem(root.left, elem)
	} else {
		return BTreeSearchItem(root.right, elem)
	}
}
