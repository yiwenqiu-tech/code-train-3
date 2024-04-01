package greedy

import "testing"

func Test_minimumEffort(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				tasks: [][]int{
					{1, 2},
					{2, 4},
					{4, 8},
				},
			},
			want: 8,
		},
		{
			name: "test2",
			args: args{
				tasks: [][]int{
					{1, 3},
					{2, 4},
					{10, 11},
					{10, 12},
					{8, 9},
				},
			},
			want: 32,
		},
		{
			name: "test3",
			args: args{
				tasks: [][]int{
					{1, 8},
					{6, 7},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffort(tt.args.tasks); got != tt.want {
				t.Errorf("minimumEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}
