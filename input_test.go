package dijkstra

import "testing"

func TestBlock2chars(t *testing.T) {
	inputc := make(chan []byte)
	retc := block2chars(inputc)
	go func() {
		inputc <- []byte("hello")
	} ()

	ret := <-retc
	if ret != 'h' {
		t.Errorf("Error: block2chars did not return correct first character.")
	}
}
