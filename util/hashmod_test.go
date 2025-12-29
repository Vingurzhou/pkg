package util_test

import (
	"testing"

	"github.com/Vingurzhou/pkg/util"
)

func TestHashMod(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key  string
		mod  uint32
		want uint32
	}{struct {
		name string
		key  string
		mod  uint32
		want uint32
	}{
		name: "test1",
		key:  "19952429930",
		mod:  16,
		want: 6,
	}, struct {
		name string
		key  string
		mod  uint32
		want uint32
	}{
		name: "test2",
		key:  "1391532161",
		mod:  16,
		want: 7,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.HashMod(tt.key, tt.mod)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("HashMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
