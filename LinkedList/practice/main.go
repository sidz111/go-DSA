// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// func Show(head *Node) {
// 	temp := head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("nil")
// }

// //func for Insert at beginning
// func InsertAtBeginning(head *Node, data int) *Node {
// 	newNode := &Node{data: data}
// 	newNode.next = head
// 	return newNode
// }

// func main() {

// 	var head *Node
// 	node1 := &Node{data: 30}
// 	node2 := &Node{data: 40}
// 	node3 := &Node{data: 50}

// 	head = node1
// 	node1.next = node2
// 	node2.next = node3

// 	newNode := &Node{data: 20}
// 	newNode.next = head
// 	head = newNode
// 	head = InsertAtBeginning(head, 2)
// 	Show(head)

// }

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

func InsertAtHead(head *Node, data int) *Node {
	newNode := &Node{data: data}
	newNode.next = head
	return newNode
}

func InsertAtTail(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		return newNode
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	return head
}

func main() {

	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}

	var head *Node
	head = node1
	node1.next = node2
	node2.next = node3
	Show(head)
	head = InsertAtHead(head, 5)
	Show(head)
	head = InsertAtTail(head, 40)
	Show(head)

	// fmt.Println(head.data)
	// fmt.Println(head.next.data)
	// fmt.Println(head.next.next.data)
}
