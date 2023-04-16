// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sorthelper_test

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"

	. "github.com/weiwenchen2022/sorthelper"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestSortIntSlice(t *testing.T) {
	t.Parallel()

	data := ints
	a := IntSlice[int]{data[:]}
	a.Sort()

	if !a.IsSorted() {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestSortFloat64Slice(t *testing.T) {
	t.Parallel()

	data := float64s
	a := Float64Slice[float64]{data[:]}
	a.Sort()

	if !a.IsSorted() {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestSortStringSlice(t *testing.T) {
	t.Parallel()

	data := strings
	a := StringSlice[string]{data[:]}
	a.Sort()

	if !a.IsSorted() {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestInts(t *testing.T) {
	t.Parallel()

	data := ints
	Ints(data[:])
	if !IntsAreSorted(data[:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestFloats(t *testing.T) {
	t.Parallel()

	data := float64s
	Float64s(data[:])
	if !Float64sAreSorted(data[:]) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestStrings(t *testing.T) {
	t.Parallel()

	data := strings
	Strings(data[:])
	if !StringsAreSorted(data[:]) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	data := strings

	SliceSort(data[:])
	if !SliceIsSorted(data[:]) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestSortLarge_Random(t *testing.T) {
	t.Parallel()

	n := 1000000
	if testing.Short() {
		n /= 100
	}
	data := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range data {
		data[i] = r.Intn(100)
	}

	if IntsAreSorted(data) {
		t.Fatalf("terrible rand.Rand")
	}

	Ints(data)
	if !IntsAreSorted(data) {
		t.Errorf("Sort didn't sort - 1M ints")
	}
}

func TestReverseSortIntSlice(t *testing.T) {
	t.Parallel()

	data := ints
	data1 := ints

	a := IntSlice[int]{data[:]}
	a.Sort()
	r := IntSlice[int]{data1[:]}
	r.Reverse()
	for i := 0; i < len(data); i++ {
		if data[i] != data1[len(data)-1-i] {
			t.Errorf("reverse sort didn't sort")
		}

		if i > len(data)/2 {
			break
		}
	}
}

type bench[E any] struct {
	name string
	f    func([]E)
}

func BenchmarkSotrtString1K(b *testing.B) {
	for _, bench := range [...]bench[string]{
		{"sort.Strings", sort.Strings},
		{"Strings", Strings[string]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			unsorted := make([]string, 1<<10)
			for i := range unsorted {
				unsorted[i] = strconv.Itoa(i ^ 0x2cc)
			}
			data := make([]string, len(unsorted))

			for i := 0; i < b.N; i++ {
				copy(data, unsorted)
				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortString1K_Slice(b *testing.B) {
	for _, bench := range [...]bench[string]{
		{
			name: "sort.Slice",
			f: func(data []string) {
				sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
			},
		},
		{
			name: "Sort",
			f:    SliceSort[string],
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			unsorted := make([]string, 1<<10)
			for i := range unsorted {
				unsorted[i] = strconv.Itoa(i ^ 0x2cc)
			}
			data := make([]string, len(unsorted))

			for i := 0; i < b.N; i++ {
				copy(data, unsorted)
				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkStableString1K(b *testing.B) {
	for _, bench := range [...]bench[string]{
		{
			name: "sort.Stable.StringSlice",
			f: func(data []string) {
				sort.Stable(sort.StringSlice(data))
			},
		},
		{
			name: "StringSlice.Stable",
			f: func(data []string) {
				StringSlice[string]{data}.Stable()
			},
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			unsorted := make([]string, 1<<10)
			for i := range unsorted {
				unsorted[i] = strconv.Itoa(i ^ 0x2cc)
			}
			data := make([]string, len(unsorted))

			for i := 0; i < b.N; i++ {
				copy(data, unsorted)
				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt1K(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{"sort.Ints", sort.Ints},
		{"Ints", Ints[int]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := range data {
					data[i] = i ^ 0x2cc
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt1K_Sorted(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{"sort.Ints", sort.Ints},
		{"Ints", Ints[int]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := range data {
					data[i] = i
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt1K_Reversed(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{"sort.Ints", sort.Ints},
		{"Ints", Ints[int]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := range data {
					data[i] = len(data) - i
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt1K_Mod8(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{"sort.Ints", sort.Ints},
		{"Ints", Ints[int]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := range data {
					data[i] = i % 8
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkStableInt1K(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{
			name: "sort.Stable.IntSlice",
			f: func(data []int) {
				sort.Stable(sort.IntSlice(data))
			},
		},
		{
			name: "IntSlice.Stable",
			f: func(data []int) {
				IntSlice[int]{data}.Stable()
			}},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			unsorted := make([]int, 1<<10)
			for i := range unsorted {
				unsorted[i] = i ^ 0x2cc
			}
			data := make([]int, len(unsorted))

			for i := 0; i < b.N; i++ {
				copy(data, unsorted)
				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkStableInt1K_Slice(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{
			name: "sort.SliceStable",
			f: func(data []int) {
				sort.SliceStable(data, func(i, j int) bool { return data[i] < data[j] })
			},
		},
		{
			name: "SliceStable",
			f:    SliceStable[int],
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			unsorted := make([]int, 1<<10)
			for i := range unsorted {
				unsorted[i] = i ^ 0x2cc
			}
			data := make([]int, len(unsorted))

			for i := 0; i < b.N; i++ {
				copy(data, unsorted)
				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt64K(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{"sort.Ints", sort.Ints},
		{"Ints", Ints[int]},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<16)
				for i := range data {
					data[i] = i ^ 0xcccc
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortInt64K_Slice(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{
			name: "sort.Slice",
			f: func(data []int) {
				sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
			},
		},
		{
			name: "Sort",
			f:    SliceSort[int],
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<16)
				for i := range data {
					data[i] = i ^ 0xcccc
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkStableInt64K(b *testing.B) {
	for _, bench := range [...]bench[int]{
		{
			name: "sort.Stable.IntSlice",
			f: func(data []int) {
				sort.Stable(sort.IntSlice(data))
			},
		},
		{
			name: "IntSlice.Stable",
			f: func(data []int) {
				IntSlice[int]{data}.Stable()
			},
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<16)
				for i := range data {
					data[i] = i ^ 0xcccc
				}

				b.StartTimer()
				bench.f(data)
				b.StopTimer()
			}
		})
	}
}

func TestStableInts(t *testing.T) {
	t.Parallel()

	data := ints
	IntSlice[int]{data[:]}.Stable()
	if !IntsAreSorted(data[:]) {
		t.Errorf("nsorted %v\n   got %v", ints, data)
	}
}
