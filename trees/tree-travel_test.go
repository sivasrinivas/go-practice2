package trees

import "testing"

func TestInOrderTravel(t *testing.T) {
	root := NewNode(5)
	root.right = NewNode(3)
	root.left = NewNode(1)

	root.right.right = NewNode(8)
	root.right.left = NewNode(2)

	root.left.right = NewNode(9)
	root.left.left= NewNode(7)

	type args struct {
		root *Node
	}
	tests := []struct {
		name        string
		args        args
		wantInOrder string
	}{
		{name: "success case", args:args{root}, wantInOrder:"7 1 9 5 2 3 8 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInOrder := InOrderTravel(tt.args.root); gotInOrder != tt.wantInOrder {
				t.Errorf("InOrderTravel() = %v, want %v", gotInOrder, tt.wantInOrder)
			}
		})
	}
}
