package data_structures

import "testing"

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	myLinkedList := LinkedList{}
	
	myLinkedList.Add(1)
	myLinkedList.Add(2)
	myLinkedList.Add(3)

	for i := 1; i <= 3; i++ {
		if myLinkedList.Search(i) == nil {
			t.Errorf("Should have contained a node with value %d", i)
		}
	}
}
