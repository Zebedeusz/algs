package exercises

import ds "algs/data_structures"

// RemoveDups removes duplicate elements from a singly linked list.
func RemoveDups(list *ds.LinkedList) {
	// option 1
	// iterate in quadratic manner through the list
	// get 1st elem, go through the list to find dups, then get 2nd, go through the list and so on
	// O(n*n)

	// option 2
	// with a map
	// iterate through the list, put all elems in the map, if map already contains one, remove from the list
	// O(n)

	// option 1 implementation

	if list == nil || list.First == nil {
		return
	}

	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		for compNode := &currNode.Next; *compNode != nil; {
			if currNode.Value == (*compNode).Value {
				if (*compNode).Next != nil {
					*compNode = (*compNode).Next
					continue
				} else {
					*compNode = nil
					break
				}
			}
			compNode = &(*compNode).Next
		}
	}
}

// KthToLastElement returns slice of values of provided linked list starting from element at index/place k to the last element.
func KthToLastElement(list *ds.LinkedList, k int) []int {
	if list == nil || list.First == nil {
		return []int{}
	}
	if k <= 0 {
		return ds.LinkedListToSlice(list)
	}

	listLen := ds.Len(list)
	if listLen < k {
		return []int{}
	}

	n := 1
	kthNodeFound := false
	listKthToLast := make([]int, 0, listLen-k)
	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		if n == k {
			kthNodeFound = true
		}
		if kthNodeFound {
			listKthToLast = append(listKthToLast, currNode.Value)
		}
		n++
	}

	return listKthToLast
}

// DeleteMiddleNode removes middle or first node from the list given only that node.
// It does not work for the last node.
func DeleteMiddleNode(node *ds.LinkedListNode) {
	if node == nil {
		return
	}
	*node = *node.Next
}

// Partition moves elements in the linked list so that elements lower than x are to left of elements greater or equal to it.
// e.g. x = 4, list = 5,7,3,4,8,1,2; result example -> 3,1,2,5,7,4,8
func Partition(list *ds.LinkedList, x int) {
	// option 1:
	// initialize new linked list
	// iterate over provided list to find elements lesser than x, add them to the new list
	// iterate again over provided list to find elements greater or equal to x, add them to the new list
	// O(n)

	if list == nil || list.First == nil {
		return
	}

	var newList *ds.LinkedList
	var lastNode *ds.LinkedListNode
	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		if currNode.Value < x {
			if newList == nil {
				newList = &ds.LinkedList{
					First: &ds.LinkedListNode{
						Value: currNode.Value,
					},
				}
				lastNode = newList.First
			} else {
				lastNode.Next = &ds.LinkedListNode{
					Value: currNode.Value,
				}
				lastNode = lastNode.Next
			}
		}
	}
	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		if currNode.Value >= x {
			if newList == nil {
				newList = &ds.LinkedList{
					First: &ds.LinkedListNode{
						Value: currNode.Value,
					},
				}
				lastNode = newList.First
			} else {
				lastNode.Next = &ds.LinkedListNode{
					Value: currNode.Value,
				}
				lastNode = lastNode.Next
			}
		}
	}
	*list = *newList
}
