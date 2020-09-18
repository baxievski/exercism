package binarysearchtree

// SearchTreeData represents a node in a binary tree
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst generates a binary search tree node
func Bst(d int) *SearchTreeData {
	return &SearchTreeData{data: d}

}

// Insert puts a new node in the binary search tree
func (b *SearchTreeData) Insert(v int) {
	if v <= b.data {
		if b.left == nil {
			b.left = Bst(v)
			return
		}
		b.left.Insert(v)
	}
	if v > b.data {
		if b.right == nil {
			b.right = Bst(v)
			return
		}
		b.right.Insert(v)
	}
}

// MapString gives a string representation of the tree
func (b *SearchTreeData) MapString(f func(int) string) []string {
	repr := []string{}
	if b == nil {
		return repr
	}
	repr = append(repr, b.left.MapString(f)...)
	repr = append(repr, f(b.data))
	repr = append(repr, b.right.MapString(f)...)
	return repr
}

// MapInt gives a integer representation of the tree
func (b *SearchTreeData) MapInt(f func(int) int) []int {
	repr := []int{}
	if b == nil {
		return repr
	}
	repr = append(repr, b.left.MapInt(f)...)
	repr = append(repr, f(b.data))
	repr = append(repr, b.right.MapInt(f)...)
	return repr
}
