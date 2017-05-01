package sim

import "fmt"

type Simulator struct {
	events *PriorityQueue
	time   float64
}

func New() *Simulator {
	s := &Simulator{}
	s.events = NewPriorityQueue()
	return s
}

func (s *Simulator) Insert(e Event) {
	s.events.Insert(e)
}

func (s *Simulator) Run() {
	for s.events.Len() > 0 {
		e := s.events.Pop()
		s.time = e.Priority()
		fmt.Println("Sim time:", s.time)
		e.Run(s)
	}
}
