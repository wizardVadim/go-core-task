package main

import (
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	if wg == nil {
		t.Fatal("NewWaitGroup returned nil")
	}
	if wg.sem == nil {
		t.Fatal("new wait group channel is nil")
	}
	if wg.count != 0 {
		t.Fatalf("count = %d, want 0", wg.count)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		initial   int
		input     int
		want      int
		wantError error
	}{
		{name: "Simple add", input: 10, want: 10},
		{name: "Negative input", input: -5, wantError: ErrNegativeCount},
		{name: "Zero"},
		{name: "Accumulate", initial: 2, input: 3, want: 5},
		{name: "Zero preserves count", initial: 2, want: 2},
		{name: "Negative preserves count", initial: 2, input: -5, want: 2, wantError: ErrNegativeCount},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := NewWaitGroup()
			if err := wg.Add(tt.initial); err != nil {
				t.Fatalf("initial Add: %v", err)
			}
			if err := wg.Add(tt.input); !errors.Is(err, tt.wantError) {
				t.Fatalf("Add error = %v, want %v", err, tt.wantError)
			}
			if wg.count != tt.want {
				t.Errorf("count = %d, want %d", wg.count, tt.want)
			}
		})
	}
}

func TestDone(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wg := NewWaitGroup()
		if err := wg.Add(1); err != nil {
			t.Fatal(err)
		}
		returned := make(chan struct{})
		go func() {
			wg.Done()
			close(returned)
		}()

		select {
		case <-wg.sem:
		case <-time.After(3 * time.Second):
			t.Fatal("Done did not send a signal")
		}
		awaitFinished(t, returned)
	})
}

func awaitFinished(t *testing.T, finished <-chan struct{}) {
	t.Helper()
	select {
	case <-finished:
	case <-time.After(3 * time.Second):
		t.Fatal("operation did not finish")
	}
}

func TestWait(t *testing.T) {
	for _, tt := range []struct {
		name  string
		count int
	}{
		{name: "Zero", count: 0},
		{name: "One task", count: 1},
		{name: "Multiple tasks", count: 10},
	} {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				wg := NewWaitGroup()
				checkWaitCycle(t, wg, tt.count)
			})
		})
	}
}

func TestWaitReuse(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wg := NewWaitGroup()
		for _, count := range []int{2, 0, 3} {
			checkWaitCycle(t, wg, count)
		}
	})
}

func checkWaitCycle(t *testing.T, wg *WaitGroup, count int) {
	t.Helper()
	if err := wg.Add(count); err != nil {
		t.Fatal(err)
	}
	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	for i := 0; i < count; i++ {
		synctest.Wait()
		select {
		case <-finished:
			t.Fatalf("Wait returned after %d of %d tasks", i, count)
		default:
		}
		go wg.Done()
	}

	awaitFinished(t, finished)
	if wg.count != 0 {
		t.Fatalf("count after Wait = %d, want 0", wg.count)
	}
}
