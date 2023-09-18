package exercises

import (
	ds "algs/data_structures"
	"strconv"
	"strings"
)

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

// SumReversedLists executes a sum of numbers stored in linked lists.
// Each number is stored in a reverse order, each node contains only one digit.
// Example: 159 = 9 -> 5 -> 1
// The result is a number stored in a linked list in the same order.
func SumReversedLists(list1, list2 *ds.LinkedList) (*ds.LinkedList, error) {
	// Possible other solution:
	// iterate over both lists at the same time
	// sum values at nodes that are pointed to, save the v%10 for the next addition operation
	// if at the end there's a value from module operation left, make it the first node of the new list

	if list1 == nil || list2 == nil || list1.First == nil || list2.First == nil {
		return nil, nil
	}

	n1, err := reversedLinkedListToNumber(list1)
	if err != nil {
		return nil, err
	}
	n2, err := reversedLinkedListToNumber(list2)
	if err != nil {
		return nil, err
	}
	sum := n1 + n2

	newList := &ds.LinkedList{}
	var lastNode *ds.LinkedListNode

	sumStr := strconv.Itoa(sum)
	sumLen := len(sumStr)
	sumRunes := []rune(sumStr)
	for i := range sumRunes {
		v, err := strconv.Atoi(string(sumRunes[sumLen-i-1]))
		if err != nil {
			return nil, err
		}

		if newList.First == nil {
			newList.First = &ds.LinkedListNode{
				Value: int(v),
			}
			lastNode = newList.First
		} else {
			lastNode.Next = &ds.LinkedListNode{
				Value: int(v),
			}
			lastNode = lastNode.Next
		}
	}

	return newList, nil
}

// Other solution:
// iterate over both lists at the same time
// sum values at nodes that are pointed to, save the v%10 for the next addition operation
// if at the end there's a value from module operation left, make it the first node of the new list
func SumReversedListsV2(list1, list2 *ds.LinkedList) *ds.LinkedList {
	if list1 == nil || list2 == nil {
		return nil
	}

	leftForNextSum := 0
	var newListLastNode *ds.LinkedListNode
	var newList *ds.LinkedList
	ptrList1 := list1.First
	ptrList2 := list2.First
	for {
		if ptrList1 == nil && ptrList2 == nil {
			break
		}
		sum := nodeValueOrZero(ptrList1) + nodeValueOrZero(ptrList2) + leftForNextSum

		leftForNextSum = sum / 10
		if sum > 9 {
			sum = sum % 10
		}

		if newListLastNode == nil {
			newListLastNode = &ds.LinkedListNode{
				Value: sum,
			}
			newList = &ds.LinkedList{
				First: newListLastNode,
			}
		} else {
			newListLastNode.Next = &ds.LinkedListNode{
				Value: sum,
			}
			newListLastNode = newListLastNode.Next
		}

		ptrList1 = ptrToNextOrNil(ptrList1)
		ptrList2 = ptrToNextOrNil(ptrList2)
	}

	return newList
}

func ptrToNextOrNil(node *ds.LinkedListNode) *ds.LinkedListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func nodeValueOrZero(node *ds.LinkedListNode) int {
	if node != nil {
		return node.Value
	}
	return 0
}

func reversedLinkedListToNumber(list *ds.LinkedList) (int, error) {
	listSlice := ds.LinkedListToSlice(list)
	listLen := len(listSlice)

	sbuilder := strings.Builder{}
	sbuilder.Grow(listLen)

	for i := range listSlice {
		if _, err := sbuilder.WriteString(strconv.Itoa(listSlice[listLen-i-1])); err != nil {
			return -1, err
		}
	}

	return strconv.Atoi(sbuilder.String())
}

// IsPalindrome checks whether a linked list is a palindrome.
// e.g. 1 -> 2 -> 3 -> 2 -> 1 will return true,
// 1 -> 0 -> 2 -> 0 will return false.
func IsPalindrome(list *ds.LinkedList) bool {
	if list == nil || list.First == nil {
		return false
	}

	listSlice := ds.LinkedListToSlice(list)
	listSliceLen := len(listSlice)

	for i := range listSlice {
		if listSlice[i] != listSlice[listSliceLen-i-1] {
			return false
		}
	}
	return true
}
