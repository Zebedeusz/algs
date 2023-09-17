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
