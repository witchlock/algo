package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {

	testTable := map[string]struct {
		Args   []int
		Result []int
	}{
		"single":    {Args: []int{1}, Result: []int{1}},
		"2":         {Args: []int{1, 2}, Result: []int{1, 2}},
		"2-reverse": {Args: []int{2, 1}, Result: []int{1, 2}},
		"odds":      {Args: []int{11, 21, 20}, Result: []int{11, 20, 21}},
		"long":      {Args: []int{2, 1, 8, 10, 70, 5, 9, 11, 80}, Result: []int{1, 2, 5, 8, 9, 10, 11, 70, 80}},
		"unequal":   {Args: []int{72, 2, 1, 8, 10, 70, 5, 9, 11, 80}, Result: []int{1, 2, 5, 8, 9, 10, 11, 70, 72, 80}},
	}

	for name, tc := range testTable {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.Result, MergeSort(tc.Args), "match")
		})
	}
}

func benchmarkMergeSort(a []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(a)
	}
}

func BenchmarkMergeSortNormal(b *testing.B) {
	a := []int{72, 2, 1, 8, 10, 70, 5, 9, 11, 80, 90, 80, 81, 78, 54, 1422, 9983}
	benchmarkMergeSort(a, b)
}

func BenchmarkMergeSortSmall(b *testing.B) {
	a := []int{72, 2, 1}
	benchmarkMergeSort(a, b)
}
