package main

import (
	"slices"
	"testing"
	"time"
)

func TestChannelFromChannels(t *testing.T) {
	tests := []struct {
		name          string
		channelsCount int
		iterations    int
		want          []int
	}{
		{
			name:          "test 5 iteration in 5 channels",
			channelsCount: 5,
			iterations:    5,
			want: []int{
				0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
			},
		},
		{
			name:          "test 5 iteration in 1 channel",
			channelsCount: 1,
			iterations:    5,
			want: []int{
				0, 1, 2, 3, 4,
			},
		},
		{
			name:          "test 1 iteration in 5 channels",
			channelsCount: 5,
			iterations:    1,
			want: []int{
				0, 0, 0, 0, 0,
			},
		},
		{
			name:          "test 0 iteration in 5 channels",
			channelsCount: 5,
			iterations:    0,
			want:          []int{},
		},
		{
			name:          "test 5 iteration in 0 channels",
			channelsCount: 0,
			iterations:    5,
			want:          []int{},
		},
		{
			name:          "test 0 iteration in 0 channels",
			channelsCount: 0,
			iterations:    0,
			want:          []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channels := make([]<-chan int, 0, tt.channelsCount)
			for range tt.channelsCount {
				ch := make(chan int)
				channels = append(channels, ch)
				go func() {
					for i := range tt.iterations {
						ch <- i
					}
					close(ch)
				}()
			}

			res := channelFromChannels(channels)
			slice := make([]int, 0, tt.channelsCount*tt.iterations)

		readLoop:
			for {
				select {
				case val, ok := <-res:
					if !ok {
						break readLoop
					}
					slice = append(slice, val)
				case <-time.After(time.Second * 20):
					t.Fatalf("no numbers from bitGenerator")
				}
			}
			slices.Sort(slice)
			slices.Sort(tt.want)
			if !slices.Equal(slice, tt.want) {
				t.Errorf("TestChannelFromChannels result slice is not equal to new; got = %+v, want = %+v", slice, tt.want)
			}
		})
	}
}
