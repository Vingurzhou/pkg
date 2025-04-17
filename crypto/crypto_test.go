package crypto

import "testing"

func TestEncryptSymmetrically(t *testing.T) {
	type args struct {
		input string
		key   rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				input: "C02DX16KQ05P",
				key:   1,
			},
			want: "B13EY07JP14Q",
		},
		{
			name: "",
			args: args{
				input: "B13EY07JP14Q",
				key:   1,
			},
			want: "C02DX16KQ05P",
		}, {
			name: "",
			args: args{
				input: "C02DX16KQ05P",
				key:   1,
			},
			want: "B13EY07JP14Q",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptSymmetrically(tt.args.input, tt.args.key); got != tt.want {
				t.Errorf("EncryptSymmetrically() = %v, want %v", got, tt.want)
			}
		})
	}
}
