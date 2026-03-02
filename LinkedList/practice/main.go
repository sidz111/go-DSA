package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func Show(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

//func for Insert at beginning
func InsertAtBeginning(head *Node, data int) *Node {
	newNode := &Node{data: data}
	newNode.next = head
	return newNode
}

func main() {

	var head *Node
	node1 := &Node{data: 30}
	node2 := &Node{data: 40}
	node3 := &Node{data: 50}

	head = node1
	node1.next = node2
	node2.next = node3

	newNode := &Node{data: 20}
	newNode.next = head
	head = newNode
	head = InsertAtBeginning(head, 2)
	Show(head)

}
