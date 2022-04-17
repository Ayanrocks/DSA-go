package Heap

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyHeap(t *testing.T) {
	ast := assert.New(t)

	mH := new(myHeap)

	heap.Init(mH)

	heap.Push(mH, 1)
	heap.Pop(mH)

	begin, end := 0, 10
	for i := begin; i < end; i++ {
		heap.Push(mH, i)
		ast.Equal(0, (*mH)[0], "The minimum value after inserting %d is %d, ih=%v", i, (*mH)[0], (*mH))
	}

	for i := begin; i < end; i++ {
		fmt.Println(i, *mH)
		ast.Equal(i, heap.Pop(mH), "Pop ih= %v", (*mH))
	}
}
