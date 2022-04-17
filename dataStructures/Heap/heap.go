package Heap

type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return item
}
