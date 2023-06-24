package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
			return root
		}
		return BTreeInsertData(root.Left, data)
	}

	if root.Right == nil {
		root.Right = &TreeNode{Data: data, Parent: root}
		return root
	}
	return BTreeInsertData(root.Right, data)
}
