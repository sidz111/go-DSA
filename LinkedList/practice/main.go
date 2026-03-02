package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func Show(head *Node) {
	temp := head //[data,next] = [4, address of node2]
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

func main() {

	node1 := &Node{data: 4}
	node2 := &Node{data: 9}
	node3 := &Node{data: 40}

	var head *Node
	head = node1
	node1.next = node2
	node2.next = node3
	Show(head)
}
