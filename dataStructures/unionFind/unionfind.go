package unionfind

type UnionFind struct {
	// Size of elements in the union find
	Size int
	// Component Size
	SZ []int

	// ID[i] points to the parent of the i, if id[i] = i then it's a root node
	ID []int

	// Total number of components
	NumComponents int
}

func New(size int) UnionFind {
	if size == 0 {
		return UnionFind{}
	}

	sz := make([]int, size)
	id := make([]int, size)

	for i := 0; i < size; i++ {
		sz[i] = 1 // Link to itself
		id[i] = i // Link to itself
	}

	return UnionFind{
		Size:          size,
		NumComponents: size,
		ID:            id,
		SZ:            sz,
	}
}

func (u *UnionFind) Find(p int) int {
	root := p

	// Finding the root of the element
	for u.ID[root] != root {
		root = u.ID[root]
	}

	// Compress the path leading back to the root
	// Doing this operation is called 'path compression'
	// this is what gives us amortized time complexity
	for p != root {
		idx := u.ID[p]
		u.ID[p] = root
		p = idx
	}

	return root
}

// returns whether p and q are connected to the same root
func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// returns the size of the component
func (u *UnionFind) ComponentSize(p int) int {
	return u.SZ[u.Find(p)]
}

func (u *UnionFind) Unify(p, q int) {

	// check if they are already connected
	if u.Connected(p, q) {
		return
	}

	root1 := u.Find(p)
	root2 := u.Find(q)

	// Merge the smaller into the larger
	if u.SZ[root1] < u.SZ[root2] {
		u.SZ[root2] += u.SZ[root1]
		u.ID[root1] = root2
		u.SZ[root1] = 0
	} else {
		u.SZ[root1] += u.SZ[root2]
		u.ID[root2] = root1
		u.SZ[root1] = 0
	}

	// reduce the total number of components by 1
	u.NumComponents--
}
