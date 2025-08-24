package main

type Node struct {
	Value int64
	Depth int64
	Right *Node
	Left  *Node
}

func CreateTree(value int64) *Node {
	tree := &Node{}
	tree.Left = nil
	tree.Right = nil
	tree.Value = value
	return tree
}

func (t *Node) Insert(value int64) {
	t.AddNode(t, value)
}

func (t *Node) AddNode(next *Node, value int64) {
	if value >= t.Value {
		if t.Right == nil {
			depth := t.Depth + 1
			t.Right = &Node{Value: value, Depth: depth}
			return
		} else {
			t.Right.AddNode(t.Right, value)
		}
	} else {
		if t.Left == nil {
			depth := t.Depth + 1
			t.Left = &Node{Value: value, Depth: depth}
			return
		} else {
			t.Left.AddNode(t.Left, value)
		}
	}

}

func Print(current *Node) {
	if current.Right != nil {
		Print(current.Right)
	}
	if current.Left != nil {
		Print(current.Left)
	}
}

func main() {
	omit_zero()
}
