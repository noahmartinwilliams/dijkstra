package dijkstra

import "sync"
func node(retc chan []string, endName string, name string) chan robot {
	inputc := make(chan robot, 100)
	go func() {
	isInf := true
	pathLength := 0.0
	path := []string{}

	for bot := range(inputc) {
		if isInf {
			isInf = false
			path = bot.path
			pathLength = bot.pathLength
		}

		if bot.pathLength < pathLength {
			path = bot.path
			pathLength = bot.pathLength
		}
	}
	retc <- path

	} ()
	return inputc
}

func treeNode(dests map[string][]dest, wg *sync.WaitGroup, name string, nodePool map[string]chan robot) chan robot {
	inputc := make(chan robot)
	go func() {
	nodec := nodePool[name]

	defer wg.Done()

	destinations, ok := dests[name]
	if !ok {
		input := <-inputc
		input.path = append(input.path, name)
		nodec <- input
	} else {
		links := make([]link, 0)
		delete(dests, name)
		for x := 0 ; x < len(destinations) ; x++ {
			wg.Add(1)
			destc := treeNode(dests, wg, destinations[x].dest, nodePool)
			links = append(links, link{dest:destc, pathLength:destinations[x].pathLength})
		}
		input := <-inputc
		go func() {
			nodec <- input
		} ()
		input.path = append(input.path, name)
		for x := 0 ; x < len(links) ; x++ {
			links[x].dest <- robot{path:input.path, pathLength:links[x].pathLength + input.pathLength}
		}
	}

	} ()
	return inputc
}
