package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	var num uint32 = 0b11111111111111111111111111111101
	for idx := 0; idx < b.N; idx++ {
		reverseBits(num)
	}
}
func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "n = 00000010100101000001111010011100",
			args: args{num: 0b00000010100101000001111010011100},
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "n = 11111111111111111111111111111101",
			args: args{num: 0b11111111111111111111111111111101},
			want: 0b10111111111111111111111111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
