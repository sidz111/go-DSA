package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// print list
func printList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

func main() {

	// nodes create karu
	n1 := &Node{data: 10}
	n2 := &Node{data: 20}
	n3 := &Node{data: 30}

	// connect karu
	n1.next = n2
	n2.next = n3

	head := n1

	printList(head)
}
