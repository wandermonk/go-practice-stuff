package main

import (
	"errors"
	"fmt"
	"log"
)

//Node represents the node element of the BST
type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

//Insert adds a value to the BST
func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("cannot add value to a nil node")
	}
	switch {
	case value == n.Value:
		return nil
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Data: data, Value: value}
			return nil
		}
		return n.Right.Insert(value, data)
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Data: data, Value: value}
			return nil
		}
		return n.Left.Insert(value, data)
	}

	return nil
}

//Find searches for the input string from the BST.
func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax(n)
}

func (n *Node) replace(parent, replacement *Node) error {
	if n == nil {
		return errors.New("cannot replace value for a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}

	parent.Right = replacement
	return nil
}

//Delete removes element from a tree
func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("cannot delete string from a nil node")
	}

	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replace(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replace(parent, n.Right)
		}

		if n.Right == nil {
			n.replace(parent, n.Left)
		}

		replacement, repParent := n.Left.findMax(n)
		n.Value = replacement.Value
		n.Data = replacement.Data
		return replacement.Delete(replacement.Value, repParent)
	}

}

//Tree the BST with nodes as left and right child
type Tree struct {
	Root *Node
}

//Insert adds a node to the tree
func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

//Find searches for a node matching s
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}

	return t.Root.Find(s)
}

//Delete removes node from tree with value matching s
func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return errors.New("Delete cannot work on nil")
	}

	fakeParent := &Node{Right: t.Root}
	err := fakeParent.Delete(s, fakeParent)

	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}

	return nil
}

//Traverse traverses through tree
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {

	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}
