package stack

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "33+22+11",
			},
			want: 66,
		},
		{
			name: "test2",
			args: args{
				s: "33+22*1",
			},
			want: 55,
		},
		{
			name: "test4",
			args: args{
				s: "3/2",
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				s: "3+5 / 2",
			},
			want: 5,
		},
		{
			name: "test6",
			args: args{
				s: "1*2-3/4+5*6-7*8+9/10",
			},
			want: -24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
