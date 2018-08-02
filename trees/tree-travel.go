package trees

import "strconv"

func InOrderTravel(root *Node) (inOrder string) {
	inOrderTravelUtil(root, &inOrder)
	return
}

func inOrderTravelUtil(root *Node, path *string) {
	if root == nil {
		return
	}

	inOrderTravelUtil(root.left, path)
	*path = *path + strconv.Itoa(root.val) + " "
	inOrderTravelUtil(root.right, path)
}

func PreOrderTravel(root *Node) (inOrder string) {
	preOrderTravelUtil(root, &inOrder)
	return
}

func preOrderTravelUtil(root *Node, path *string) {
	if root == nil {
		return
	}

	*path = *path + strconv.Itoa(root.val) + " "
	preOrderTravelUtil(root.left, path)
	preOrderTravelUtil(root.right, path)
}

func PostOrderTravel(root *Node) (inOrder string) {
	postOrderTravelUtil(root, &inOrder)
	return
}

func postOrderTravelUtil(root *Node, path *string) {
	if root == nil {
		return
	}

	postOrderTravelUtil(root.left, path)
	postOrderTravelUtil(root.right, path)
	*path = *path + strconv.Itoa(root.val) + " "
}
