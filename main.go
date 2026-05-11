package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (t *LinkedList) AddNode(value int) {
	newNode := &Node{Value: value, Next: nil}

	if t.Head == nil {
		t.Head = newNode
		return
	}
	current := t.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (t *LinkedList) PrintList() {
	current := t.Head

	for current != nil {
		fmt.Printf("%d ->", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (t *LinkedList) DeleteNode(value int) {
	if t.Head == nil {
		return
	}
	if t.Head.Value == value {
		t.Head = t.Head.Next
		return
	}
	current := t.Head

	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (t *LinkedList) InsertNodeAt(value int, position int) *Node {
	newNode := &Node{Value: value, Next: nil}

	if position == 0 {
		newNode.Next = t.Head
		t.Head = newNode
		return nil
	}
	current := t.Head

	for i := 0; i < position-1; i++ {
		if current == nil {
			fmt.Println("Error to return add the node")
		}
		current = current.Next
	}
	if current == nil {
		fmt.Println("Error to return add the node")
	}
	newNode.Next = current.Next
	current.Next = newNode
	return nil
}

func main() {
	var (
		op    int
		vl    int
		list  LinkedList
		index int
	)
	for {
		ListMenu()
		_, _ = fmt.Scanln(&op)

		switch op {
		case 0:
			return
		case 1:
			fmt.Println("What valor do you want to ADD?:")
			_, _ = fmt.Scanln(&vl)
			list.AddNode(vl)
			screen.Clear()
		case 2:
			list.PrintList()
		case 3:
			fmt.Println("What valor do you want to DELETE?:")
			_, _ = fmt.Scanln(&vl)
			list.AddNode(vl)
			screen.Clear()
		case 4:
			fmt.Println("What valor do you want to add, and in which index?:")
			_, _ = fmt.Scanln(&vl, &index)
			list.InsertNodeAt(vl, index)
		}
	}
}
