package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.data)
		current = current.next
		fmt.Print(" ")
	}

}

//1>2>3>4

//1>2>3>4

func printReversList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev

}

func main() {

	node5 := &Node{data: 5, next: nil}
	node4 := &Node{data: 4, next: node5}
	node3 := &Node{data: 3, next: node4}
	node2 := &Node{data: 2, next: node3}
	node1 := &Node{data: 1, next: node2}

	printList(node1)
	fmt.Println(" ")
	reverseList := printReversList(node1)
	printList(reverseList)

}
