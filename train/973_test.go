package train

import (
	"reflect"
	"testing"
)

func Test_kClosest2(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				points: [][]int{
					{0, 1},
					{1, 0},
				},
				k: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest2(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
