package bm1342

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name string
		nums int
		want int
	}{
		{
			name: "Example 1",
			nums: 14,
			want: 6,
		},
		{
			name: "Example 2",
			nums: 8,
			want: 4,
		},
		{
			name: "Example 3",
			nums: 123,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfSteps(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
