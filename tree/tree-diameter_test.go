package tree

import "testing"

func Test_treeDiameter(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				edges: [][]int{{0, 1}, {1, 2}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeDiameter(tt.args.edges); got != tt.want {
				t.Errorf("treeDiameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
