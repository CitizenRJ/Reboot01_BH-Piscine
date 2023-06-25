package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		left := BTreeInsertData(root.Left, data)
		root.Left = left
		left.Parent = root
	} else {
		right := BTreeInsertData(root.Right, data)
		root.Right = right
		right.Parent = root
	}
	return root
}
