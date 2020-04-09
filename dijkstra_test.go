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

	mpc := launchNodes(retc, endName, inputc)
	mp := <-mpc
	_, ok := mp["a"]
	if !ok {
		t.Errorf("Error: launchNodes failed to launch first node.")
	}
}

func TestLaunchTreeNodes(t *testing.T) {
	inputc := make(chan dest)
	go func() {
		defer close(inputc)
		inputc <- dest{source:"a", dest:"b", pathLength:1}
		inputc <- dest{source:"b", dest:"c", pathLength:2}
	} ()
	str := launchTreeNodes("c", "a", inputc)
	if str[2] != "c" {
		t.Errorf("Error: launchTreeNodes did not return proper end node.")
	}
}
