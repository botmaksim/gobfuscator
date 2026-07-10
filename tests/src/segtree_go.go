package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type SegmentTree struct {
	tree []int64
	n    int
}

func NewSegmentTree(size int) *SegmentTree {
	return &SegmentTree{
		tree: make([]int64, 4*size),
		n:    size,
	}
}

func (st *SegmentTree) build(a []int64, v, tl, tr int) {
	if tl == tr {
		st.tree[v] = a[tl]
	} else {
		tm := (tl + tr) / 2
		st.build(a, v*2, tl, tm)
		st.build(a, v*2+1, tm+1, tr)
		st.tree[v] = st.tree[v*2] + st.tree[v*2+1]
	}
}

func (st *SegmentTree) query(v, tl, tr, l, r int) int64 {
	if l > r {
		return 0
	}
	if l == tl && r == tr {
		return st.tree[v]
	}
	tm := (tl + tr) / 2
	return st.query(v*2, tl, tm, l, min(r, tm)) +
		st.query(v*2+1, tm+1, tr, max(l, tm+1), r)
}

func (st *SegmentTree) update(v, tl, tr, pos int, new_val int64) {
	if tl == tr {
		st.tree[v] = new_val
	} else {
		tm := (tl + tr) / 2
		if pos <= tm {
			st.update(v*2, tl, tm, pos, new_val)
		} else {
			st.update(v*2+1, tm+1, tr, pos, new_val)
		}
		st.tree[v] = st.tree[v*2] + st.tree[v*2+1]
	}
}

func main() {
	n := 5
	a := []int64{1, 2, 3, 4, 5}
	st := NewSegmentTree(n)
	st.build(a, 1, 0, n-1)

	fmt.Printf("Sum 0-4: %d\n", st.query(1, 0, n-1, 0, 4))
	st.update(1, 0, n-1, 2, 10)
	fmt.Printf("Sum 0-4 after update: %d\n", st.query(1, 0, n-1, 0, 4))

	msg := "Segment tree runs perfectly!"
	fmt.Println(msg)
}
