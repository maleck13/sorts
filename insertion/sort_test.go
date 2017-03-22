package insertion_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/sorts/insertion"
)

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		TestName string
		Arr      []int
		Expected string
	}{
		{
			TestName: "simple",
			Arr:      []int{2, 7, 6, 4, 5, 3, 1},
			Expected: "[1 2 3 4 5 6 7]",
		},
		{
			TestName: "duplicate",
			Arr:      []int{2, 2, 7, 6, 4, 5, 3, 1},
			Expected: "[1 2 2 3 4 5 6 7]",
		},
		{
			TestName: "reverse",
			Arr:      []int{7, 6, 5, 4, 3, 2, 1, 0},
			Expected: "[0 1 2 3 4 5 6 7]",
		},
	}

	for _, tc := range cases {
		t.Run(tc.TestName, func(t *testing.T) {
			insertion.Sort(tc.Arr)
			out := fmt.Sprintf("%v", tc.Arr)
			if out != tc.Expected {
				t.Fatal("arr was not sorted expected " + tc.Expected + " but got " + out)
			}
		})
	}

}
