package a1672

import (
	"reflect"
	"testing"
)

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		name   string
		values [][]int
		want   int
	}{
		{
			name:   "Test 1",
			values: [][]int{{1, 2, 3}, {3, 2, 1}},
			want:   6,
		},
		{
			name:   "Test 2",
			values: [][]int{{1, 5}, {7, 3}, {3, 5}},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumWealth(tt.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
