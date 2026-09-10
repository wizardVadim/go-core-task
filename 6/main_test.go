package main

import (
	"testing"
	"time"
)

func TestRandomIntGenerator(t *testing.T) {
	tests := []struct {
		name string
		bits []int
		max  int
		want int
	}{
		{
			name: "Example",
			bits: []int{1, 1, 1, 1},
			max:  10,
			want: 5,
		},
		{
			name: "1010 becomes zero after modulo 10",
			bits: []int{1, 0, 1, 0},
			max:  10,
			want: 0,
		},
		{
			name: "1111 becomes five after modulo 10",
			bits: []int{1, 1, 1, 1},
			max:  10,
			want: 5,
		},
		{
			name: "maximum is one",
			bits: []int{1},
			max:  1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			go func() {
				for _, v := range tt.bits {
					ch <- v
				}
			}()
			result := getRandomInt(ch, tt.max)
			if result != tt.want {
				t.Errorf("TestRandomIntGenerator result unexpectable. Got = %d, want = %d", result, tt.want)
			}
		})
	}
}

func TestBitGenerator(t *testing.T) {
	bitCh := make(chan int)

	go bitGenerator(bitCh)

	testSize := 1000
	for range testSize {
		select {
		case bit := <-bitCh:
			if bit != 0 && bit != 1 {
				t.Errorf("generator returned %d instead of 1 or 0", bit)
			}
		case <-time.After(time.Second * 20):
			t.Fatalf("no numbers from bitGenerator")
		}
	}
}
