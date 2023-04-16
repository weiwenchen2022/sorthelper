// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sorthelper_test

import (
	"fmt"

	"sorthelper"
)

// A Change is a record of source code changes, recording user, language, and delta size.
type Change struct {
	user     string
	language string
	lines    int
}

var changes = []Change{
	{"gri", "Go", 100},
	{"ken", "C", 150},
	{"glenda", "Go", 200},
	{"rsc", "Go", 200},
	{"r", "Go", 100},
	{"ken", "Go", 200},
	{"dmr", "C", 100},
	{"r", "C", 150},
	{"gri", "Smalltalk", 80},
}

// ExampleMultiKeys demonstrates sorting a struct type using different sets of multiple fields in the comparison.
// We chain together "Less" functions with MultiSorter, each of which compares a single field.
func Example_sortMultiKeys() {
	// Closures that order the Change structure.
	user := func(c1, c2 *Change) bool {
		return c1.user < c2.user
	}
	language := func(c1, c2 *Change) bool {
		return c1.language < c2.language
	}
	increasingLines := func(c1, c2 *Change) bool {
		return c1.lines < c2.lines
	}
	decreasingLines := func(c1, c2 *Change) bool {
		return c1.lines > c2.lines // Note: > orders downwards.
	}

	// Simple use: Sort by user.
	sorthelper.NewMultiSorter(changes).OrderedBy(user)
	fmt.Println("By user:", changes)

	// More examples.
	sorthelper.NewMultiSorter(changes).OrderedBy(user, increasingLines)
	fmt.Println("By user,<lines:", changes)

	sorthelper.NewMultiSorter(changes).OrderedBy(user, decreasingLines)
	fmt.Println("By user,>lines:", changes)

	sorthelper.NewMultiSorter(changes).OrderedBy(language, increasingLines)
	fmt.Println("By language,<lines:", changes)

	sorthelper.NewMultiSorter(changes).OrderedBy(language, increasingLines, user)
	fmt.Println("By language,<lines,user:", changes)

	// Output:
	// By user: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
	// By user,<lines: [{dmr C 100} {glenda Go 200} {gri Smalltalk 80} {gri Go 100} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
	// By user,>lines: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken Go 200} {ken C 150} {r C 150} {r Go 100} {rsc Go 200}]
	// By language,<lines: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
	// By language,<lines,user: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
}
