package dijkstra

func launchNodes(retc chan []string, endName string, inputc chan dest) map[string]chan robot {
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
	return ret
}
