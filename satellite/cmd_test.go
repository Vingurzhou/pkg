package satellite

import "testing"

func TestScrambleStrToStr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试用例1",
			args: args{
				input: "C2 34 54 10 00 1B 00 00 00 25 00 01 02 03 04 05 00 01 02 03 11 00 11 00 11 00",
			},
			want:    "3d0dca4a68f206f56cac2fa0335d0cc552a9b9ad5fc2d6ed77dc",
			wantErr: false,
		},
		{
			name: "测试用例2",
			args: args{
				input: "3d0dca4a68f206f56cac2fa0335d0cc552a9b9ad5fc2d6ed77dc",
			},
			want:    "c2345410001b0000002500010203040500010203110011001100",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScrambleHexString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScrambleStrToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ScrambleStrToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
