package util_test

import (
	"testing"

	"github.com/Vingurzhou/pkg/util"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr []int
	}{{
		name: "",
		arr:  []int{1, 3, 5, 7, 11},
	}, {
		name: "",
		arr:  []int{5, 3, 1, 7, 11},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			util.BubbleSort(tt.arr)
		})
	}
}
