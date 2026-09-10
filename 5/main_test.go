package main

import (
	"slices"
	"testing"
)

func TestOriginSliceElements(t *testing.T) {
	tests := []struct {
		name     string
		origin   []int
		other    []int
		wantBool bool
		want     []int
	}{
		{
			name:     "Example",
			origin:   []int{65, 3, 58, 678, 64},
			other:    []int{64, 2, 3, 43},
			wantBool: true,
			want:     []int{3, 64},
		},
		{
			name:     "No matches",
			origin:   []int{65, 3, 58, 678, 64},
			other:    []int{5, 10},
			wantBool: false,
			want:     []int{},
		},
		{
			name:     "All elements match",
			origin:   []int{1, 2},
			other:    []int{2, 1},
			wantBool: true,
			want:     []int{1, 2},
		},
		{
			name:     "Other is longer",
			origin:   []int{64, 2, 3, 43},
			other:    []int{65, 3, 58, 678, 64},
			wantBool: true,
			want:     []int{64, 3},
		},
		{
			name:     "Duplicates",
			origin:   []int{65, 3, 65, 3},
			other:    []int{3, 3},
			wantBool: true,
			want:     []int{3, 3},
		},
		{name: "Empty origin", origin: []int{}, other: []int{1}, wantBool: false, want: []int{}},
		{name: "Nil origin", origin: nil, other: []int{1}, wantBool: false, want: []int{}},
		{name: "Empty other", origin: []int{63, 5}, other: []int{}, wantBool: false, want: []int{}},
		{name: "Nil other", origin: []int{65, 3}, other: nil, wantBool: false, want: []int{}},
		{name: "Both nil", origin: nil, other: nil, wantBool: false, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originBefore := slices.Clone(tt.origin)
			otherBefore := slices.Clone(tt.other)

			ok, got := isAcrossSlices(tt.origin, tt.other)
			if ok != tt.wantBool {
				t.Errorf("isAcrossSlices() isAcross = %v, want %v", ok, tt.wantBool)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("isAcrossSlices() = %v, want %v", got, tt.want)
			}
			if !slices.Equal(tt.origin, originBefore) {
				t.Errorf("origin changed: got %v, want %v", tt.origin, originBefore)
			}
			if !slices.Equal(tt.other, otherBefore) {
				t.Errorf("other changed: got %v, want %v", tt.other, otherBefore)
			}
		})
	}
}
