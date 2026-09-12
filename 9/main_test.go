package main

import (
	"slices"
	"testing"
)

func TestConvertNums(t *testing.T) {
	tests := []struct {
		name  string
		input []uint8
		want  []float64
	}{
		{
			name:  "Empty input",
			input: []uint8{},
			want:  []float64{},
		},
		{
			name:  "Nil input",
			input: nil,
			want:  []float64{},
		},
		{
			name:  "Zero",
			input: []uint8{0},
			want:  []float64{0},
		},
		{
			name:  "One",
			input: []uint8{1},
			want:  []float64{1},
		},
		{
			name:  "Maximum uint8",
			input: []uint8{255},
			want:  []float64{16581375},
		},
		{
			name:  "Multiple numbers preserve order and duplicates",
			input: []uint8{4, 13, 11, 6, 44, 13, 11, 67},
			want:  []float64{64, 2197, 1331, 216, 85184, 2197, 1331, 300763},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch1 := make(chan uint8)
			ch2 := make(chan float64)

			go func() {
				for _, v := range tt.input {
					ch1 <- v
				}
				close(ch1)
			}()

			go convertNums(ch1, ch2)

			got := make([]float64, 0, len(tt.input))
			for v := range ch2 {
				got = append(got, v)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("test convert nums; unexpectable slice; got = %+v, want = %+v", got, tt.want)
			}
		})
	}
}
