// This file wraps binary search.

package sorthelper

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Search searches for x in a sorted slice of a and returns the index.
// The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
func Search[E constraints.Ordered](a []E, x E) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// Convenience wrappers for common cases.

// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchInts[E constraints.Integer](a []E, x E) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// SearchFloat64s searches for x in a sorted slice of float64s and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchFloat64s[E ~float64](a []E, x E) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// SearchStrings searches for x in a sorted slice of strings and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchStrings[E ~string](a []E, x E) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// Search returns the result of applying SearchInts to the receiver and x.
func (p IntSlice[E]) Search(x E) int { return SearchInts(p.Slice, x) }

// Search returns the result of applying SearchFloat64s to the receiver and x.
func (p Float64Slice[E]) Search(x E) int { return SearchFloat64s(p.Slice, x) }

// Search returns the result of applying SearchStrings to the receiver and x.
func (p StringSlice[E]) Search(x E) int { return SearchStrings(p.Slice, x) }
