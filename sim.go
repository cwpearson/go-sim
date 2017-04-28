package sim

import "fmt"

type Simulator struct {
	events *EventQueue
}

func New() *Simulator {
	s := &Simulator{}
	s.events = NewEventQueue()
	return s
}

func (s *Simulator) Insert(e Event) {
	s.events.Push(e)
}

func (s *Simulator) Run() {
	for s.events.Size() > 0 {
		e := s.events.Pop().(Event)
		priority := e.Priority()
		fmt.Println(priority)
		e.Run()
	}
}
