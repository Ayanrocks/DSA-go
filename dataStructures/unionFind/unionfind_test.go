package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnionFind(t *testing.T) {
	ast := assert.New(t)

	// Creating an element of 10 sizes
	u := New(6)

	// ID:		[0,1,2,3,4,5]
	// Index: 	[0,1,2,3,4,5]
	// Size: 	[1,1,1,1,1,1]

	u.Unify(1, 2)
	u.Unify(0, 3)
	u.Unify(4, 5)

	// ID:		[0,1,1,0,4,4]
	// Index: 	[0,1,2,3,4,5]
	// Size: 	[1,1,1,1,1,1]

	ast.Equal(u.ID[0], u.ID[3], "Should be true")
	ast.Equal(u.ID[1], u.ID[2], "Should be true")
	ast.Equal(u.ID[4], u.ID[5], "Should be true")

	u.Unify(0, 1)
	u.Unify(3, 4)

	// ID:		[0,0,1,0,0,4]
	// Index: 	[0,1,2,3,4,5]
	// Size: 	[1,1,1,1,1,1]

	ast.Equal(u.ID[0], u.ID[1], "Should be true")
	ast.Equal(u.ID[3], u.ID[4], "Should be true")

	u.Find(2)

	ast.Equal(u.ID[0], u.ID[2], "Should do path compression")
}
