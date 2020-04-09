package dijkstra

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
