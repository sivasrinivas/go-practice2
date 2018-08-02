package trees

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{val: val}
}
