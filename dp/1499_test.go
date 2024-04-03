package dp

import "testing"

func Test_findMaxValueOfEquation(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				points: [][]int{
					{1, 3},
					{2, 0},
					{5, 10},
					{6, -10},
				},
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxValueOfEquation(tt.args.points, tt.args.k); got != tt.want {
				t.Errorf("findMaxValueOfEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
