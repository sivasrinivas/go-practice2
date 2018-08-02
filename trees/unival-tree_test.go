package trees

import "testing"

func TestCountUnivalTrees(t *testing.T) {

	threeNodeTreeRoot := NewNode(5)
	threeNodeTreeRoot.right = NewNode(5)
	threeNodeTreeRoot.left = NewNode(5)

	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "single node tree", args: struct{ root *Node }{root: NewNode(5)}, want: 0},
		{name: "three node tree", args: args{threeNodeTreeRoot}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountUnivalTrees(tt.args.root); got != tt.want {
				t.Errorf("CountUnivalTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
