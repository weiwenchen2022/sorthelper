// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sorthelper_test

import (
	"fmt"

	"sorthelper"
)

// This example demonstrates searching a list sorted in ascending order.
func ExampleSearch() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sorthelper.Search(a, x)
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	// Output:
	// found 6 at index 2 in [1 3 6 10 15 21 28 36 45 55]
}

// This example demonstrates searching for float64 in a list sorted in ascending order.
func ExampleSearchFloat64s() {
	type myfloat64 float64
	a := []myfloat64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	x := myfloat64(2.0)
	i := sorthelper.SearchFloat64s(a, x)
	fmt.Printf("found %g at index %d in %v\n", x, i, a)

	x = myfloat64(0.5)
	i = sorthelper.SearchFloat64s(a, x)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", x, i, a)

	// Output:
	// found 2 at index 1 in [1 2 3.3 4.6 6.1 7.2 8]
	// 0.5 not found, can be inserted at index 0 in [1 2 3.3 4.6 6.1 7.2 8]
}

// This example demonstrates searching for int in a list sorted in ascending order.
func ExampleSearchInts() {
	type myint int
	a := []myint{1, 2, 3, 4, 6, 7, 8}

	x := myint(2)
	i := sorthelper.SearchInts(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = myint(5)
	i = sorthelper.SearchInts(a, x)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)

	// Output:
	// found 2 at index 1 in [1 2 3 4 6 7 8]
	// 5 not found, can be inserted at index 4 in [1 2 3 4 6 7 8]
}
