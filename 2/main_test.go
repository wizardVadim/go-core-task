package main

import (
	"slices"
	"testing"
)

func TestRandomIntSlice(t *testing.T) {
	tests := []struct {
		name        string
		inputLength int
		maxValue    int64
	}{
		{
			name:        "1",
			inputLength: 10,
			maxValue:    100,
		},
		{
			name:        "2",
			inputLength: 15,
			maxValue:    10000,
		},
		{
			name:        "3",
			inputLength: 40,
			maxValue:    10,
		},
		{
			name:        "Empty slice",
			inputLength: 0,
			maxValue:    10,
		},
		{
			name:        "Single element",
			inputLength: 1,
			maxValue:    100,
		},
		{
			name:        "Only zeros",
			inputLength: 10,
			maxValue:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice, err := randomIntSlice(tt.inputLength, tt.maxValue)
			if err != nil {
				t.Fatalf("TestRandomIntSlice, error = %v", err)
			}

			length := len(slice)
			if length != tt.inputLength {
				t.Errorf("incorrect slice length; got = %d, want = %d", length, tt.inputLength)
			}

			for i, v := range slice {
				if v < 0 || v >= tt.maxValue {
					t.Errorf("slice[%d] = %d, want 0 <= value < %d",
						i, v, tt.maxValue)
				}
			}
		})
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "Best way",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:  []int64{0, 2, 4, 6, 8},
		},
		{
			name:  "No even nums",
			input: []int64{1, 3, 3, 5, 7},
			want:  []int64{},
		},
		{
			name:  "All even nums",
			input: []int64{0, 2, 4, 6},
			want:  []int64{0, 2, 4, 6},
		},
		{
			name:  "Empty slice",
			input: []int64{},
			want:  []int64{},
		},
		{
			name:  "One Even value",
			input: []int64{2},
			want:  []int64{2},
		},
		{
			name:  "One Odd value",
			input: []int64{1},
			want:  []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exampled := sliceExample(tt.input)
			if slices.Compare(exampled, tt.want) != 0 {
				t.Errorf("TestSliceExample got = %v, want = %v", exampled, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		value int64
		want  []int64
	}{
		{
			name:  "Best way",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 10,
			want:  []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "Zero value",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 0,
			want:  []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			name:  "Empty origin slice",
			input: []int64{},
			value: 10,
			want:  []int64{10},
		},
		{
			name:  "Non-init origin slice",
			input: nil,
			value: 10,
			want:  []int64{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := addElements(tt.input, tt.value)
			if slices.Compare(out, tt.want) != 0 {
				t.Errorf("TestAddElements got = %v, want = %v", out, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "Best way",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:  []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Empty origin slice",
			input: []int64{},
			want:  []int64{},
		},
		{
			name:  "Non-init origin slice",
			input: nil,
			want:  []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := copySlice(tt.input)
			if slices.Compare(out, tt.want) != 0 {
				t.Errorf("TestAddElements got = %v, want = %v", out, tt.want)
			}
			if len(out) > 0 {
				before := tt.input[0]
				out[0] ^= 1

				if tt.input[0] != before {
					t.Error("TestAddElements same slices after copy")
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		idx   int
		want  []int64
	}{
		{
			name:  "Best way",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			idx:   3,
			want:  []int64{0, 1, 2, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Last element",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			idx:   9,
			want:  []int64{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "First element",
			input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			idx:   0,
			want:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Only element",
			input: []int64{42},
			idx:   0,
			want:  []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := removeElement(tt.input, tt.idx)
			if slices.Compare(out, tt.want) != 0 {
				t.Errorf("TestAddElements got = %v, want = %v", out, tt.want)
			}
		})
	}
}
