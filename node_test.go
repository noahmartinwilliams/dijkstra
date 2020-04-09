package dijkstra

import "testing"

func TestNode(t *testing.T) {
	retc := make(chan []string)
	inputc := node(retc, "e", "e")

	go func() {
		inputc <- robot{path:[]string{"s", "e"}, pathLength:1}
		close(inputc)
	} ()

	ret := <-retc
	if ret[0] != "s" {
		t.Errorf("Error: node did not return correct start name in first test.")
	}

}
