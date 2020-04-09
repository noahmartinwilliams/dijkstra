package dijkstra

func block2chars(inputc chan []byte) chan byte {
	retc := make(chan byte, 1024)
	go func() {

	defer close(retc)
	for input := range(inputc) {
		for x := 0 ; x < len(input) ; x++ {
			retc <- input[x]
		}
	}

	} ()
	return retc
}
