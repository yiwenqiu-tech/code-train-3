package search

import "testing"

func Test_estimateCost(t *testing.T) {
	type args struct {
		items *[2][3]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				items: &[2][3]int{{4, 2, 3}, {1, 5, 0}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := estimateCost(tt.args.items); got != tt.want {
				t.Errorf("estimateCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
