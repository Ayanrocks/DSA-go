package priorityqueue

import "container/heap"

type Item struct {
	Name  string
	Value int // this denotes the priority

	// Index is the Index of the entry in the heap
	// this is used to determine the position of the element without searching in the entire tree
	Index int
}

type PriorityQueue []*Item

// retunrns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// this is the comparator function which compares two elements in priority queue
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

// fetches the root of the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	// Copies the priority queue
	old := *pq
	n := len(old)

	// gets the last element of the priority queue
	item := old[n-1]

	// sets the index of the last element as -1
	item.Index = -1

	// Overwrites the queue with the copy by removing the last element
	*pq = old[0 : n-1]

	// returns the last element
	return item
}

// Swaps two elements in the queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

	pq[i].Index = i
	pq[j].Index = j
}

// Push pushes an element to the end of the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	element := x.(*Item)
	element.Index = len(*pq)
	*pq = append(*pq, element)
}

func (pq *PriorityQueue) update(item *Item, name string, value int) {
	item.Name = name
	item.Value = value
	heap.Fix(pq, item.Index)
}
