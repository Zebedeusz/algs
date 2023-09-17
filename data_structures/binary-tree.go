package datastructures

import (
	"fmt"
	"time"
)

type Node struct {
	Left  *Node
	Right *Node
	value int
}

type Tree struct {
	Root *Node
}

func Insert(tree *Tree, value int) {
	if tree.Root == nil {
		tree.Root = &Node{
			value: value,
		}
		return
	}

	toCheck := []*Node{tree.Root}
	for {
		nodeToCheck := toCheck[0]

		if nodeToCheck.Left == nil {
			nodeToCheck.Left = &Node{
				value: value,
			}
			return
		}
		if nodeToCheck.Right == nil {
			nodeToCheck.Right = &Node{
				value: value,
			}
			return
		}

		toCheck = toCheck[1:]
		toCheck = append(toCheck, nodeToCheck.Left, nodeToCheck.Right)
	}
}

func FindSubTree(node *Node, value int) *Tree {
	if node == nil {
		return nil
	}
	if node.value == value {
		return &Tree{Root: node}
	}
	if found := FindSubTree(node.Left, value); found != nil {
		return found
	}
	if found := FindSubTree(node.Right, value); found != nil {
		return found
	}
	return nil
}

func Find(tree *Tree, value int) *Node {
	if tree.Root == nil {
		return nil
	}
	if tree.Root.value == value {
		return tree.Root
	}
	if found := Find(&Tree{Root: tree.Root.Left}, value); found != nil {
		return found
	}
	if found := Find(&Tree{Root: tree.Root.Right}, value); found != nil {
		return found
	}
	return nil
}

func Delete(tree *Tree, value int) {
	toCheck := []**Node{&tree.Root}
	ticker := time.NewTicker(time.Second)

	for {
		nodeToCheck := toCheck[0]

		emptyNode := &Node{}
		var toMerge **Node = &(emptyNode)

		if nodeToCheck == nil || *nodeToCheck == nil {
			return
		}
		if (*nodeToCheck).value == value {
			if (*nodeToCheck).Left != nil {
				*toMerge = (*nodeToCheck).Right
				*nodeToCheck = (*nodeToCheck).Left
			} else if (*nodeToCheck).Right != nil {
				*nodeToCheck = (*nodeToCheck).Right
			} else {
				*nodeToCheck = nil
			}

			// insert nodes to the right into the tree
			if pointsToNotEmptyNode(toMerge) {
				nodesToMergeChan := make(chan *Node)
				go WalkNodesLevelOrder(&Tree{Root: *toMerge}, nodesToMergeChan)
				for {
					select {
					case v := <-nodesToMergeChan:
						Insert(tree, v.value)
					case <-ticker.C:
						ticker.Stop()
						return
					}
				}
			}

			return
		}

		toCheck = toCheck[1:]
		toCheck = append(toCheck, &(*nodeToCheck).Left, &(*nodeToCheck).Right)
	}
}

func pointsToNotEmptyNode(nodePtr **Node) bool {
	nodePtrPtr := *nodePtr
	return nodePtr != nil &&
		nodePtrPtr != nil &&
		(nodePtrPtr.value != 0 ||
			nodePtrPtr.Left != nil ||
			nodePtrPtr.Right != nil)
}

func Walk(t *Node, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

//	 1
//	 /\
//	2   3
// /\  /\
// 4 5 6 7
// /
// 8

// -> 1, 2, 3, 4, 5, 6, 7, 8
func WalkNodesLevelOrder(t *Tree, ch chan *Node) {
	ch <- t.Root

	nodesToCheck := []*Node{t.Root}
	for {
		if len(nodesToCheck) == 0 {
			return
		}

		nodeToCheck := nodesToCheck[0]

		if nodeToCheck.Left != nil {
			ch <- nodeToCheck.Left
		}
		if nodeToCheck.Right != nil {
			ch <- nodeToCheck.Right
		}

		nodesToCheck = nodesToCheck[1:]
		if nodeToCheck.Left != nil {
			nodesToCheck = append(nodesToCheck, nodeToCheck.Left)
		}
		if nodeToCheck.Right != nil {
			nodesToCheck = append(nodesToCheck, nodeToCheck.Right)
		}
	}
}

// GetTreeValues returns tree values ordered from bottom-left of the tree, in the direction of going up
func GetTreeValues(node *Node) []int {
	values := make([]int, 0, 10)

	if node == nil {
		return []int{}
	}

	values = append(values, GetTreeValues(node.Left)...)
	values = append(values, node.value)
	values = append(values, GetTreeValues(node.Right)...)

	return values
}

func PrintTree(root *Node) {
	if root == nil {
		return
	}
	PrintTree(root.Left)
	fmt.Println(root.value)
	PrintTree(root.Right)
}

func CopyValues(src, dst *Tree) {
	valuesChan := make(chan int)
	ticker := time.NewTicker(time.Second)

	go Walk(src.Root, valuesChan)
	for {
		select {
		case value := <-valuesChan:
			Insert(dst, value)
		case <-ticker.C:
			ticker.Stop()
			return
		}
	}
}

func DeepCopy(src, dst *Tree) {
	nodesChan := make(chan *Node)
	ticker := time.NewTicker(time.Second)

	go WalkNodesLevelOrder(src, nodesChan)
	for {
		select {
		case node := <-nodesChan:
			Insert(dst, node.value)
		case <-ticker.C:
			ticker.Stop()
			return
		}
	}
}
