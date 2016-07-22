package data_structures

// Linked List Implementation
type LinkedList struct{
	first *Node
	last *Node
	length int
}

type Node struct{
	element int
	next *Node
	prev *Node
}

func (this *LinkedList) addElement(elem int){
	anotherNode := Node{element : elem, next : nil, prev: this.last}

	if this.first == nil {
		this.first = &anotherNode
	} else {
		this.last = &anotherNode	
	}
	this.length += 1
}

func (this *LinkedList) searchElement(elem int) *Node{
	for node := this.first; node != nil; node = node.next {
		if node.element == elem {
			return node
		}
	}
	return nil
}

var linkedList = LinkedList{first: nil, last: nil, length: 0}

func Add(elem int) {
	linkedList.addElement(elem)
	return
}

func Search(elem int) *Node {
	return linkedList.searchElement(elem)
}
