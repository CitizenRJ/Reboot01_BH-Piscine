package piscine

//type TreeNode struct {
//	Data  string
//	Left  *TreeNode
//	Right *TreeNode
//}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	} else if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}
