package util

import "testing"

func TestContainsString(t *testing.T) {
	type args struct {
		slice []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{struct {
		name string
		args args
		want bool
	}{
		name: "",
		args: args{
			slice: []string{"1", "2"},
			str:   "1",
		},
		want: true,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.slice, tt.args.str); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
