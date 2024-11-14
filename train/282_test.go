package train

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				num:    "1000000009",
				target: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v, len:%v", got, tt.want, len(got))
			}
		})
	}
}

func Test_calNum(t *testing.T) {
	type args struct {
		num []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				num: []byte("2+2"),
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				num: []byte("2+2+3"),
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				num: []byte("2+2*3"),
			},
			want: 8,
		},
		{
			name: "test2",
			args: args{
				num: []byte("2*2+3"),
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				num: []byte("1*2*3-4*5+6+7-8*9"),
			},
			want: -73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calNum(tt.args.num); got != tt.want {
				t.Errorf("calNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
