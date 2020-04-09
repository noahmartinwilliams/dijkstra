package dijkstra

type robot struct {
	path []string
	pathLength float64
}

type dest struct {
	dest string
	pathLength float64
}

type link struct {
	dest chan robot
	pathLength float64
}
