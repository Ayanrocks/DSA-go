package binarySearchTree

import "fmt"

type BST struct {
	Val   int
	Left  *BST
	Right *BST
}

const (
	IN_ORDER   = "IN_ORDER"
	POST_ORDER = "POST_ORDER"
	PRE_ORDER  = "PRE_ORDER"
)

// New returns a new BST
func New(val int) *BST {
	return &BST{
		Val: val,
	}
}

func (b *BST) Add(elem int) *BST {
	if b.Contains(elem) {
		return b
	}

	b = add(b, elem)
	return b
}

func add(b *BST, elem int) *BST {
	if b == nil {
		b = New(elem)
	} else if elem < b.Val {
		b.Left = add(b.Left, elem)
	} else {
		b.Right = add(b.Right, elem)
	}
	return b
}

func (b *BST) Contains(elem int) bool {
	if elem == b.Val {
		return true
	}

	if elem < b.Val && b.Left != nil {
		return b.Left.Contains(elem)
	}

	if elem > b.Val && b.Right != nil {
		return b.Right.Contains(elem)
	}

	return false

}

func (b *BST) Remove(elem int) *BST {
	if b.Contains(elem) {
		return remove(b, elem)
	}

	return b
}

func remove(b *BST, elem int) *BST {
	if b == nil {
		return b
	}

	if elem < b.Val {
		b.Left = remove(b.Left, elem)
	} else if elem > b.Val {
		b.Right = remove(b.Right, elem)
	} else {
		// check if left node is empty

		if b.Left == nil {
			return b.Right
		} else if b.Right == nil {
			return b.Left
		} else {
			tmp := b.Right.FindMin()

			b.Val = tmp.Val

			b.Right = remove(b.Right, tmp.Val)
		}
	}

	return b
}

func (b BST) FindMin() *BST {
	for b.Left != nil {
		b = *b.Left
	}

	return &b
}

func (b BST) FindMax() *BST {
	for b.Right != nil {
		b = *b.Right
	}

	return &b
}

func (b *BST) Traverse(traversalType string) {
	switch traversalType {
	case IN_ORDER:
		inOrder(b)
	case POST_ORDER:
		postOrder(b)
	case PRE_ORDER:
		preOrder(b)
	default:
		return
	}
	fmt.Println("")
}

func (b *BST) TraverseIterative(traversalType string) {
	switch traversalType {
	case IN_ORDER:
		inOrderIterative(b)
	case POST_ORDER:
		postOrderIterative(b)
	case PRE_ORDER:
		preOrderIterative(b)
	default:
		return
	}
	fmt.Println("")
}

func preOrder(b *BST) {
	if b == nil {
		return
	}

	fmt.Print(b.Val, " ")
	preOrder(b.Left)
	preOrder(b.Right)

}

func inOrder(b *BST) {
	if b == nil {
		return
	}

	inOrder(b.Left)
	fmt.Print(b.Val, " ")
	inOrder(b.Right)
}

func postOrder(b *BST) {
	if b == nil {
		return
	}

	postOrder(b.Left)
	postOrder(b.Right)
	fmt.Print(b.Val, " ")
}

func preOrderIterative(b *BST) {
	stack := []*BST{}

	stack = append(stack, b)

	for len(stack) != 0 {
		// pop
		item := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		fmt.Print(item.Val, " ")

		if item.Right != nil {
			stack = append(stack, item.Right)
		}

		if item.Left != nil {
			stack = append(stack, item.Left)
		}
	}

}

func inOrderIterative(b *BST) {
	stack := []*BST{}
	curr := b
	stack = append(stack, curr)

	for len(stack) != 0 || curr != nil {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(curr.Val, " ")

		curr = curr.Right
	}
}

func postOrderIterative(b *BST) {
	s1 := []*BST{}
	s2 := []int{}

	s1 = append(s1, b)

	for len(s1) != 0 {
		item := s1[len(s1)-1]

		s2 = append(s2, item.Val)

		s1 = append(s1, item.Right, item.Left)
	}

	for i := len(s2) - 1; i >= 0; i-- {
		fmt.Print(s2[i], " ")
	}
}
