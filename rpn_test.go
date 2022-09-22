package go_rpn

import "testing"

func TestRpn(t *testing.T) {
	type args struct {
		input []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"逆ポーランド記法の計算結果が返される",
			args{input: []interface{}{4, 6, 2, "+", "*", 3, 1, "-", 5, "*", "-"}},
			22,
			false,
		},
		{
			"無効な文字列が含まれる場合はエラーが返される",
			args{input: []interface{}{4, 6, 2, "＋", "*", 3, 1, "-", 5, "*", "-"}},
			0,
			true,
		},
		{
			"無効な値が含まれる場合はエラーが返される",
			args{input: []interface{}{4.0, 6, 2, "+", "*", 3, 1, "-", 5, "*", "-"}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rpn(tt.args.input...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rpn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rpn() got = %v, want %v", got, tt.want)
			}
		})
	}
}
