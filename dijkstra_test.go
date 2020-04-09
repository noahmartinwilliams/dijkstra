package dijkstra

import "testing"

func TestLaunchNodes(t *testing.T) {
	retc := make(chan []string)
	endName := "c"
	inputc := make(chan dest)
	go func() {
		defer close(inputc)
		inputc <- dest{source:"a", dest:"b", pathLength:1}
		inputc <- dest{source:"b", dest:"c", pathLength:2}
	} ()

	mp := launchNodes(retc, endName, inputc)
	_, ok := mp["a"]
	if !ok {
		t.Errorf("Error: launchNodes failed to launch first node.")
	}
}
