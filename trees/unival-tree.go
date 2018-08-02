package trees

import "fmt"



func main() {
	root := NewNode(5)
	root.right = NewNode(5)
	root.left = NewNode(5)

	root.right.right = NewNode(5)
	root.right.left = NewNode(5)

	root.left.right = NewNode(5)
	root.left.left= NewNode(1)

	fmt.Println(CountUnivalTrees(root))

}

func CountUnivalTrees(root *Node) int {
	var count = 0
	countUnivalTreeUtil(root, &count)
	return count
}

func countUnivalTreeUtil(root *Node, count *int) bool {
	if root == nil {
		return false
	}

	countUnivalTreeUtil(root.left, count)
	countUnivalTreeUtil(root.right, count)

	if root.right == nil || root.left == nil {
		return false
	}

	if root.left != nil && root.val != root.left.val {
		return false
	}

	if root.right != nil && root.val != root.right.val {
		return false
	}
	*count = *count + 1
	return true
}
