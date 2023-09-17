package datastructures

type LinkedListNode struct {
	Next  *LinkedListNode
	Value int
}

type LinkedList struct {
	First *LinkedListNode
}

func CreateLinkedListFromSlice(slice []int) LinkedList {
	list := LinkedList{}
	if len(slice) == 0 {
		return list
	}

	list.First = &LinkedListNode{
		Value: slice[0],
	}
	currListNode := list.First
	for _, elem := range slice[1:] {
		currListNode.Next = &LinkedListNode{
			Value: elem,
		}
		currListNode = currListNode.Next
	}
	return list
}

func LinkedListToSlice(list *LinkedList) []int {
	listArr := make([]int, 0, 10)
	if list == nil || list.First == nil {
		return listArr
	}
	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		listArr = append(listArr, currNode.Value)
	}

	return listArr
}

func Add(list *LinkedList, elem int) {
	if list == nil {
		return
	}
	if list.First == nil {
		list.First = &LinkedListNode{
			Value: elem,
		}
		return
	}

	for currNode := list.First; ; currNode = currNode.Next {
		if currNode.Next == nil {
			currNode.Next = &LinkedListNode{
				Value: elem,
			}
			return
		}
	}
}

func Remove(list *LinkedList, elem int) {
	if list == nil || list.First == nil {
		return
	}
	for currNode := &list.First; ; currNode = &(*currNode).Next {
		if (*currNode).Value == elem {
			if (*currNode).Next != nil {
				*currNode = (*currNode).Next
			} else {
				*currNode = nil
			}
			return
		}
	}
}

func Len(list *LinkedList) int {
	if list == nil || list.First == nil {
		return 0
	}

	len := 0
	for currNode := list.First; currNode != nil; currNode = currNode.Next {
		len++
	}
	return len
}
