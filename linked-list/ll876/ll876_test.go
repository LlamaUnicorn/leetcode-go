package ll876

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test1",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			name: "test2",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(&tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
