package binarySearchTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bst(t *testing.T) {
	ast := assert.New(t)
	bst := New(9)

	//    				  9
	// 			4					15
	// 		3		5			11
	//	1					10		13

	bst.Add(4)
	bst.Add(15)
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(11)
	bst.Add(10)
	bst.Add(13)

	// InOrder: 	1 3 4 5 9 10 11 13 15
	// PreOrder: 	9 4 3 1 5 15 11 10 13
	// PostOrder: 	1 3 5 4 10 13 11 15 9

	fmt.Println("Traversing Recursion")
	bst.Traverse(IN_ORDER)
	bst.Traverse(PRE_ORDER)
	bst.Traverse(POST_ORDER)
	ast.Equal(bst.Val, 9, "BST Root should be 9")

	fmt.Println("Traversing Iteratively - 1")
	bst.TraverseIterative(IN_ORDER)
	bst.TraverseIterative(PRE_ORDER)
	bst.TraverseIterative(POST_ORDER)

	bst.Remove(11)
	bst.Remove(15)

	bst.Traverse(IN_ORDER)
	bst.Traverse(PRE_ORDER)
	bst.Traverse(POST_ORDER)

	fmt.Println("Traversing Iteratively - 2")
	bst.TraverseIterative(IN_ORDER)
	bst.TraverseIterative(PRE_ORDER)
	bst.TraverseIterative(POST_ORDER)

}
