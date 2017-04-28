package sim

type LessThanComparable interface {
	Less(other LessThanComparable) bool
}

type Event interface {
	LessThanComparable
	Run()
}

type anEvent struct {
	time float64
}

func (e *anEvent) Run() {
	return
}

func (e *anEvent) Less(other *anEvent) bool {
	return e.time < other.time
}

// EventQueue implements heap.Interface and holds Items.
type EventQueue []Event

func (eq *EventQueue) Len() int { return len(*eq) }

func (pq EventQueue) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (eq EventQueue) Size() int {
	return len(eq)
}

func (eq *EventQueue) Push(x interface{}) {
	// n := eq.Len()
	e := x.(Event)
	*eq = append(*eq, e)
}

func (eq *EventQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	item := old[n-1]
	// item.index = -1 // for safety
	*eq = old[0 : n-1]
	return item
}

func (pq EventQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	// 	pq[i].index = i
	// 	pq[j].index = j
}

// // update modifies the priority and value of an Item in the queue.
// func (pq *EventQueue) update(item *Item, value string, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

// // This example creates a EventQueue with some items, adds and manipulates an item,
// // and then removes the items in priority order.
// func main() {
// 	// Some items and their priorities.
// 	items := map[string]int{
// 		"banana": 3, "apple": 2, "pear": 4,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(EventQueue, len(items))
// 	i := 0
// 	for value, priority := range items {
// 		pq[i] = &Item{
// 			value:    value,
// 			priority: priority,
// 			index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		value:    "orange",
// 		priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	pq.update(item, item.value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("%.2d:%s ", item.priority, item.value)
// 	}
// }
