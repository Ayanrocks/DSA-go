package priorityqueue

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_priorityQueue(t *testing.T) {
	ast := assert.New(t)

	// Finding servers based on their load using priority queue
	servers := map[string]int{
		"server_1": 4,
		"server_2": 2,
		"server_3": 1,
		"server_4": 3,
	}

	pq := make(PriorityQueue, len(servers))
	i := 0

	for name, value := range servers {
		pq[i] = &Item{
			Name:  name,
			Value: value,
			Index: i,
		}

		i++
	}

	// Creating a new heap for priority queue
	heap.Init(&pq)

	newServer := &Item{
		Name:  "server_5",
		Value: 6,
	}

	heap.Push(&pq, newServer)

	// Update the priority of the newly entered server to 0
	pq.update(newServer, newServer.Name, 0)

	/*
		The tree will be like this
								   	Server_5 - 0
					Server_3 - 1					Server_2 - 2
		Server_4 - 3		Server_1 - 4
	*/

	// Some items and their priorities.
	expected := []string{
		"server_5",
		"server_3",
		"server_2",
		"server_4",
		"server_1",
	}

	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		// Pop the root of the tree
		it := heap.Pop(&pq).(*Item)
		// Compare with the expected array
		ast.Equal(expected[it.Value], it.Name)
	}

}
