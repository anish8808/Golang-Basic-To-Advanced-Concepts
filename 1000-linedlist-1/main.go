package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// type Linkedlist struct {
// 	head *Node
// }

// //add linked list as last to the end

// func (ll *Linkedlist) addlist(data int) {
// 	newNode := &Node{data: data}

// 	// list is empty
// 	if ll.head == nil {
// 		ll.head = newNode
// 		return
// 	}

// 	//added in last
// 	currNode := ll.head

// 	for currNode.next != nil {
// 		currNode = currNode.next
// 	}
// 	currNode = newNode

// }

// //print the list

// func (ll *Linkedlist) printLinkedList() {
// 	if ll.head == nil {
// 		return
// 	}

// 	currNode := ll.head

// 	for currNode.next != nil {
// 		fmt.Println(currNode.data, " ")
// 		currNode = currNode.next
// 	}
// }

// func (ll *Linkedlist) addedTofirst(data int) {
// 	newNode := &Node{data: data, next: ll.head}

// 	newNode.next = ll.head

// 	ll.head = newNode
// }

// func main() {
// 	l1 := Linkedlist{}
// 	l1.addlist(12)
// 	l1.addlist(13)
// 	l1.addlist(14)
// 	l1.printLinkedList()

// }

func appendlast(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		return newNode
	}

	currNode := head
	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = newNode

	return head
}

//appending the fist

func appendFirst(head *Node, data int) *Node {
	newNode := &Node{data: data, next: head.next}

	return newNode
}

//print the List

func PrintLL(head *Node) {
	currNode := head

	for currNode != nil {
		fmt.Print(currNode.data, "->")
		currNode = currNode.next
	}
	fmt.Println("NULL")
}

func main() {

	var head *Node
	head = appendlast(head, 12)
	head = appendlast(head, 13)
	head = appendlast(head, 14)
	head = appendlast(head, 15)

	PrintLL(head)

}
