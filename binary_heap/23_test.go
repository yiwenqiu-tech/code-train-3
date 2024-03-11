package binary_heap

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				lists: []*ListNode{
					nil,
					nil,
				},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				lists: []*ListNode{
					{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
					{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
