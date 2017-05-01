package sim

import "sort"

type Event interface {
	Run(s *Simulator)
	Priority() float64
}

type ByPriority []Event

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority() < a[j].Priority() }

// EventQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	events []Event
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		events: []Event{},
	}
}

func (pq *PriorityQueue) Len() int { return len(pq.events) }

func (pq *PriorityQueue) Insert(e Event) {
	pq.events = append(pq.events, e)
	sort.Sort(ByPriority(pq.events))
}

func (pq *PriorityQueue) Pop() Event {
	e := pq.events[0]
	pq.events = pq.events[1:]
	return e
}
