package dp

import "testing"

func Test_maxProfit123(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit123(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit123() = %v, want %v", got, tt.want)
			}
		})
	}
}
