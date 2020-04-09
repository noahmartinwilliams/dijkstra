package dijkstra

import "sync"

func launchNodes(retc chan []string, endName string, inputc chan dest) chan map[string]chan robot {
	retc2 := make(chan map[string]chan robot)
	go func() {
		ret := make(map[string]chan robot)
		for input := range(inputc) {
			_, ok := ret[input.source]
			if !ok {
				ret[input.source] = node(retc, endName, input.source)
			}
			_, ok = ret[input.dest]
			if !ok {
				ret[input.dest] = node(retc, endName, input.dest)
			}

		}
		retc2 <- ret
	} ()
	return retc2
}

func splitter(inputc chan dest, out1 chan dest, out2 chan dest) {
	go func() {

	defer close(out1)
	defer close(out2)
	for input := range(inputc) {
		out1 <- input
		out2 <- input
	}

	} ()
}

func collectNodes(inputc chan dest) map[string][]dest {
	ret := make(map[string][]dest)
	for input := range(inputc) {
		list, ok := ret[input.source]
		if ok {
			ret[input.source] = append(list, input)
		} else {
			ret[input.source] = []dest{input}
		}
	}
	return ret
}

func launchTreeNodes(endName string, startName string, inputc chan dest) []string {
	retc := make(chan []string)
	inputc2 := make(chan dest)
	inputc3 := make(chan dest)
	splitter(inputc, inputc2, inputc3)

	mpc := launchNodes(retc, endName, inputc2)
	dests := collectNodes(inputc3)

	mp := <-mpc
	var wg sync.WaitGroup
	wg.Add(1)
	inc := treeNode(dests, &wg, startName, mp)
	inc <- robot{path:[]string{}, pathLength:0.0}
	wg.Wait()
	for _, value := range(mp) {
		close(value)
	}
	ret := <-retc
	return ret
}
