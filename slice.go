package sorthelper

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SliceSort sorts the slice x as determined by the operator <, in increasing order.
//
// The sort is not guaranteed to be stable: equal elements
// may be reversed from their original order.
// For a stable sort, use StableSort.
func SliceSort[E constraints.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

// SliceStable sorts the slice x using the operator <, in ascending order,
// keeping equal elements in their original order.
func SliceStable[E constraints.Ordered](x []E) {
	sort.SliceStable(x, func(i, j int) bool { return x[i] < x[j] })
}

// SliceIsSorted reports whether the slice s is sorted in increasing order according to the operator <.
func SliceIsSorted[E constraints.Ordered](x []E) bool {
	return sort.SliceIsSorted(x, func(i, j int) bool { return x[i] < x[j] })
}
