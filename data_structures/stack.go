package datastructures

type Stack interface {
	Push(v int)
	Pop() int
}

type StackLinkedList struct {
	first *LinkedListNode
}

func (s *StackLinkedList) Push(v int) {
	if s.first == nil {
		s.first = &LinkedListNode{
			Value: v,
		}
		return
	}
	newFirst := &LinkedListNode{
		Next:  s.first,
		Value: v,
	}
	s.first = newFirst
}

func (s *StackLinkedList) Pop() int {
	if s.first == nil {
		return -1
	}
	v := s.first.Value
	s.first = s.first.Next
	return v
}
