package main

import "errors"

type WaitGroup struct {
	count int
	sem   chan struct{}
}

var (
	ErrNegativeCount = errors.New("Negative task count")
)

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(n int) error {
	if n < 0 {
		return ErrNegativeCount
	}
	wg.count += n
	return nil
}

func (wg *WaitGroup) Done() {
	wg.sem <- struct{}{}
}

func (wg *WaitGroup) Wait() {
	for wg.count > 0 {
		<-wg.sem
		wg.count--
	}
}
