package main

import (
	"slices"
	"testing"
)

func TestOriginSliceElements(t *testing.T) {
	tests := []struct {
		name   string
		origin []string
		other  []string
		want   []string
	}{
		{
			name:   "Example",
			origin: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			other:  []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "No matches preserves order",
			origin: []string{"cherry", "apple", "banana"},
			other:  []string{"fig"},
			want:   []string{"cherry", "apple", "banana"},
		},
		{
			name:   "All elements excluded",
			origin: []string{"apple", "banana"},
			other:  []string{"banana", "apple"},
			want:   []string{},
		},
		{
			name:   "Other is longer",
			origin: []string{"apple"},
			other:  []string{"banana", "cherry", "fig"},
			want:   []string{"apple"},
		},
		{
			name:   "Duplicates",
			origin: []string{"apple", "banana", "apple", "banana"},
			other:  []string{"banana", "banana"},
			want:   []string{"apple", "apple"},
		},
		{
			name:   "Empty string retained",
			origin: []string{"", "apple", ""},
			other:  []string{"apple"},
			want:   []string{"", ""},
		},
		{
			name:   "Empty string excluded",
			origin: []string{"", "apple"},
			other:  []string{""},
			want:   []string{"apple"},
		},
		{name: "Empty origin", origin: []string{}, other: []string{"apple"}, want: []string{}},
		{name: "Nil origin", origin: nil, other: []string{"apple"}, want: []string{}},
		{name: "Empty other", origin: []string{"banana", "apple"}, other: []string{}, want: []string{"banana", "apple"}},
		{name: "Nil other", origin: []string{"banana", "apple"}, other: nil, want: []string{"banana", "apple"}},
		{name: "Both nil", origin: nil, other: nil, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originBefore := slices.Clone(tt.origin)
			otherBefore := slices.Clone(tt.other)

			got := originSliceElements(tt.origin, tt.other)
			if !slices.Equal(got, tt.want) {
				t.Errorf("originSliceElements() = %q, want %q", got, tt.want)
			}
			if !slices.Equal(tt.origin, originBefore) {
				t.Errorf("origin changed: got %q, want %q", tt.origin, originBefore)
			}
			if !slices.Equal(tt.other, otherBefore) {
				t.Errorf("other changed: got %q, want %q", tt.other, otherBefore)
			}
		})
	}
}
