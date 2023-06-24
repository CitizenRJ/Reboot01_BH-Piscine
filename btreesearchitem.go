package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if root.Data == elem {
			return root
		}
		if elem < root.Data {
			return BTreeSearchItem(root.Left, elem)
		} else {
			return BTreeSearchItem(root.Right, elem)
		}
	} else {
		return nil
	}
}
