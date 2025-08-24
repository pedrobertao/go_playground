package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

type Node struct {
	Value int64
	Depth int64
	Right *Node
	Left  *Node
}

func New(value int64) *Node {
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

func PrintToString(current *Node) string {
	result := fmt.Sprintf("Depth: %d, Value: %d\n", current.Depth, current.Value)

	if current.Right != nil {
		result += PrintToString(current.Right)
	}
	if current.Left != nil {
		result += PrintToString(current.Left)
	}

	return result
}

func createTree(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Error: createTree requires 1 argument"
	}

	value := int64(args[0].Int())
	tree := New(value)

	treeData, _ := json.Marshal(map[string]interface{}{
		"value": tree.Value,
		"depth": tree.Depth,
		"id":    fmt.Sprintf("tree_%d", value),
	})

	return string(treeData)
}

func insertNode(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Error: insertNode requires 2 arguments"
	}

	rootValue := int64(args[0].Int())
	insertValue := int64(args[1].Int())

	tree := New(rootValue)
	tree.Insert(insertValue)

	return PrintToString(tree)
}

func buildTree(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: buildTree requires at least 1 argument"
	}

	tree := New(int64(args[0].Int()))

	for i := 1; i < len(args); i++ {
		tree.Insert(int64(args[i].Int()))
	}

	return PrintToString(tree)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("createTree", js.FuncOf(createTree))
	js.Global().Set("insertNode", js.FuncOf(insertNode))
	js.Global().Set("buildTree", js.FuncOf(buildTree))

	fmt.Println("Go WebAssembly binary tree functions loaded")
	<-c
}
