package main

import sim "github.com/cwpearson/go-sim"

type anEvent struct {
	time float64
}

func (e *anEvent) Run() {
	return
}

func (e *anEvent) Priority() float64 {
	return e.time
}

func main() {
	s := sim.New()

	s.Insert(&anEvent{0.0})
	s.Insert(&anEvent{1.0})
	s.Insert(&anEvent{0.0})
	s.Insert(&anEvent{2.0})

	s.Run()
}
