package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

func (p *linkedList) removeNode(index int) (string, error) {

	if p.head == nil {
		return "", errors.New("Nothing to print. Linked List is empty!")
	}

	if index < 1 || index > p.size {
		return "", errors.New("Invalid index position.")
	}

	var item string

	if index == 1 {
		item = p.head.item
		p.head = p.head.next
	} else {
		currentNode := p.head
		prevNode := p.head
		for i := 1; i <= index-1; i++ {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		item = currentNode.item
		prevNode.next = currentNode.next
	}
	p.size--
	return item, nil
}

func (p *linkedList) addAtPos(index int, item string) error {
	if p.head == nil {
		return errors.New("Empty linked list.")
	}

	if index < 1 || index > p.size {
		return errors.New("Invalid index position.")
	}

	// currentNode := p.head // pushing item back to fit item
	// for i:= 1; i <= index-1; i++ {
	// 	currentNode := currentNode.next //what next?
	// }

	if index == 1 { //adding to first postition
		p.head.item = item
	} else { // adding to other positions
		currentNode := p.head
		for i := 1; i <= index-1; i++ {
			if i == index-1 {
				currentNode.next.item = item
			}
			currentNode = currentNode.next
		}
		currentNode.item = item
	}
	p.size++
	return nil
}

func (p *linkedList) get(index int) (string, error) {

	if p.head == nil {
		return "", errors.New("Empty linked list.")
	}

	if index < 1 || index > p.size {
		return "", errors.New("Invalid index position.")
	}

	var item string

	for i := 1; i <= index; i++ { // not working
		currentNode := p.head
		if i == index-1 {
			item = currentNode.item
		}
		currentNode = currentNode.next
	}

	return item, nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil //why do we need to return anything?
}

func main() {
	myList := &linkedList{nil, 0} //init link list

	myList.addNode("Peter")
	myList.addNode("Parker")
	myList.addNode("Mary Jane")
	myList.addNode("SpooderMan")
	myList.addNode("Batman")
	myList.addNode("Superman")
	myList.addNode("Wonder woman")
	myList.addNode("Flash")

	fmt.Println("\nPrinting all the nodes in the Linked List...")
	myList.printAllNodes()

	// fmt.Println("Remove one node")
	// myList.removeNode(2)

	fmt.Println("\nAdding a node at position 3...")
	myList.addAtPos(3, "Venom")
	myList.printAllNodes()

	fmt.Println("\nGetting a node at position 7..")
	fmt.Println(myList.get(6))
}
