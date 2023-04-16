// Package sorthelper provides convenience for sorting slices and user-defined collections.
package sorthelper

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

// Slice help implements sort.Interface by providing Len and
// Swap methods for embedding value.
type Slice[E any] []E

func (x Slice[E]) Len() int      { return len(x) }
func (x Slice[E]) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Convenience types for common cases

// IntSlice implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded slice value, sorting in increasing order.
type IntSlice[E constraints.Integer] struct{ Slice[E] }

func (x IntSlice[E]) Less(i, j int) bool { return x.Slice[i] < x.Slice[j] }

// Sort is a convenience method: x.Sort() calls sort.Sort(x).
func (x IntSlice[E]) Sort() { sort.Sort(x) }

// Reverse is a convenience method: x.Reverse() calls sorrt.Sort(sort.Reverse(x)).
func (x IntSlice[E]) Reverse() { sort.Sort(sort.Reverse(x)) }

// Stable is a convenience method: x.Stable() calls sort.Stable(x).
func (x IntSlice[E]) Stable() { sort.Stable(x) }

// IsSorted is a convenience method: x.IsSorted() calls sort.IsSorted(x).
func (x IntSlice[E]) IsSorted() bool { return sort.IsSorted(x) }

// Float64Slice implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded slice value, sorting in increasing order,
// with not-a-number (NaN) values ordered before other values.
type Float64Slice[E ~float64] struct{ Slice[E] }

// Less reports whether x[i] should be ordered before x[j], as required by the sort Interface.
// Note that floating-point comparison by itself is not a transitive relation: it does not
// report a consistent ordering for not-a-number (NaN) values.
// This implementation of Less places NaN values before any others, by using:
//
//	x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
func (x Float64Slice[E]) Less(i, j int) bool {
	return x.Slice[i] < x.Slice[j] || (math.IsNaN(float64(x.Slice[i])) && !math.IsNaN(float64(x.Slice[j])))
}

// Sort is a convenience method: x.Sort() calls sort.Sort(x).
func (x Float64Slice[E]) Sort() { sort.Sort(x) }

// Stable is a convenience method: x.Stable() calls sort.Stable(x).
func (x Float64Slice[E]) Stable() { sort.Stable(x) }

// Reverse is a convenience method: x.Reverse() calls sort.Sort(sort.Reverse(x)).
func (x Float64Slice[E]) Reverse() { sort.Sort(sort.Reverse(x)) }

// IsSorted is a convenience method: x.IsSorted() calls sort.IsSorted(x).
func (x Float64Slice[E]) IsSorted() bool { return sort.IsSorted(x) }

// StringSlice implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded slice value, sorting in increasing order.
type StringSlice[E ~string] struct{ Slice[E] }

func (x StringSlice[E]) Less(i, j int) bool { return x.Slice[i] < x.Slice[j] }

// Sort is a convenience method: x.Sort() calls sort.Sort(x).
func (x StringSlice[E]) Sort() { sort.Sort(x) }

// Stable is a convenience method: x.Stable() calls sort.Stable(x).
func (x StringSlice[E]) Stable() { sort.Stable(x) }

// Reverse is a convenience method: x.Reverse() calls sort.Sort(sort.Reverse(x)).
func (x StringSlice[E]) Reverse() { sort.Sort(sort.Reverse(x)) }

// IsSorted is a convenience method: x.IsSorted() calls sort.IsSorted(x).
func (x StringSlice[E]) IsSorted() bool { return sort.IsSorted(x) }

// Convenience wrappers for common cases

// Ints sorts a slice of ints in increasing order.
func Ints[E constraints.Integer](x []E) { IntSlice[E]{x}.Sort() }

// Float64s sorts a slice of floats in increasing order.
// Not-a-number (NaN) values are ordered before other values.
func Float64s[E ~float64](x []E) { Float64Slice[E]{x}.Sort() }

// Strings sorts a slice of strings in increasing order.
func Strings[E ~string](x []E) { StringSlice[E]{x}.Sort() }

// IntsAreSorted reports whether the slice x is sorted in increasing order.
func IntsAreSorted[E constraints.Integer](x []E) bool { return IntSlice[E]{x}.IsSorted() }

// Float64sAreSorted reports whether the slice s is sorted in increasing order,
// with not-a-number (NaN) values before any other values.
func Float64sAreSorted[E ~float64](x []E) bool { return Float64Slice[E]{x}.IsSorted() }

// StringsAreSorted reports whether the slice x is sorted in increasing order.
func StringsAreSorted[E ~string](x []E) bool { return StringSlice[E]{x}.IsSorted() }

// Sorter joins a by function and a slice s to be sorted.
type Sorter[E any] struct {
	s  []E
	by func(e1, e2 *E) bool // The function (closure) that defines the sort order.
}

// NewSorter returns a Sorter that sorts the slice s.
// Call its OrderedBy method to sort the slice within using the by functions.
func NewSorter[E any](s []E) *Sorter[E] {
	return &Sorter[E]{
		s: s,
	}
}

func (s *Sorter[E]) Len() int { return len(s.s) }

func (s *Sorter[E]) Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i] }

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *Sorter[E]) Less(i, j int) bool { return s.by(&s.s[i], &s.s[j]) }

// OrderedBy sorts the slice within according to the by function.
// The sort is not guaranteed to be stable. For a stable sort, use StableBy.
func (s *Sorter[E]) OrderedBy(by func(e1, e2 *E) bool) {
	s.by = by
	sort.Sort(s)
}

// StableBy sorts the slice within according to the by function,
// while keeping the original order of equal elements.
func (s *Sorter[E]) StableBy(by func(e1, e2 *E) bool) {
	s.by = by
	sort.Stable(s)
}

// MultiSorter implements the Sort interface, sorting the slice within.
type MultiSorter[E any] struct {
	s    []E
	less []func(e1, e2 *E) bool
}

// NewMultiSorter returns a MulitSorter that sorts the argument slice.
// Call its OrderedBy method to sort the slice within using the less functions, in order.
func NewMultiSorter[E any](s []E) *MultiSorter[E] {
	return &MultiSorter[E]{
		s: s,
	}
}

func (ms *MultiSorter[E]) Len() int { return len(ms.s) }

func (ms *MultiSorter[E]) Swap(i, j int) { ms.s[i], ms.s[j] = ms.s[j], ms.s[i] }

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other).
// Note that it can call the less functions twice per call.
func (ms *MultiSorter[E]) Less(i, j int) bool {
	p, q := &ms.s[i], &ms.s[j]

	// Try all but the last comparison.
	k := 0
	for ; k < len(ms.less)-1; k++ {
		less := ms.less[k]

		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}

	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

// OrderedBy sorts the slice within according to the less functions, in order.
// The sort is not guaranteed to be stable. For a stable sort, use StableBy.
func (ms *MultiSorter[E]) OrderedBy(less ...func(e1, e2 *E) bool) {
	ms.less = less
	sort.Sort(ms)
}

// StableBy sorts the slice within in ascending order as determined by the less functions, in order,
// while keeping the original order of equal elements.
func (ms *MultiSorter[E]) StableBy(less ...func(e1, e2 *E) bool) {
	ms.less = less
	sort.Stable(ms)
}
