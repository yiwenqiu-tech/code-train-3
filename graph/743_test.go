package graph

import "testing"

func Test_networkDelayTime2(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				times: [][]int{{1, 2, 1}},
				n:     2,
				k:     2,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				n:     4,
				k:     2,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				times: [][]int{{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43}, {4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0}, {2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59}},
				n:     5,
				k:     3,
			},
			want: 78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime3(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}
