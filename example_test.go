// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sorthelper_test

import (
	"fmt"
	"math"

	"sorthelper"
)

func ExampleInts() {
	type myint int
	s := []myint{5, 2, 6, 3, 1, 4} // unsorted
	sorthelper.Ints(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5 6]
}

func ExampleIntsAreSorted() {
	type myint int
	s := []myint{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sorthelper.IntsAreSorted(s))

	s = []myint{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sorthelper.IntsAreSorted(s))

	s = []myint{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sorthelper.IntsAreSorted(s))

	// Output:
	// true
	// false
	// false
}

func ExampleFloat64s() {
	type myfloat64 float64
	s := []myfloat64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sorthelper.Float64s(s)
	fmt.Println(s)

	s = []myfloat64{myfloat64(math.Inf(1)), myfloat64(math.NaN()), myfloat64(math.Inf(-1)), 0.0} // unsorted
	sorthelper.Float64s(s)
	fmt.Println(s)

	// Output:
	// [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]
}

func ExampleFloat64sAreSorted() {
	type myfloat64 float64
	s := []myfloat64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	fmt.Println(sorthelper.Float64sAreSorted(s))

	s = []myfloat64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	fmt.Println(sorthelper.Float64sAreSorted(s))

	s = []myfloat64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	fmt.Println(sorthelper.Float64sAreSorted(s))

	// Output:
	// true
	// false
	// false
}

func ExampleStrings() {
	type mystring string
	s := []mystring{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sorthelper.Strings(s)
	fmt.Println(s)
	// Output:
	// [Alpha Bravo Delta Go Gopher Grin]
}
