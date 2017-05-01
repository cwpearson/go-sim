package sim

type Wire struct {
	i, o bool
}

type And struct {
	in  [2]bool
	out [1]bool
}
