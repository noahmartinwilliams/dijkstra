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

func TestChars2lines(t *testing.T) {
	inputc := make(chan []byte)
	retc := chars2lines(block2chars(inputc))
	go func() {
		inputc <- []byte("hello\nworld")
		close(inputc)
	} ()
	ret := <-retc
	if ret != "hello" {
		t.Errorf("Error: chars2lines did not return correct first line.")
	}
	ret2 := <-retc
	if ret2 != "world" {
		t.Errorf("Error: chars2lines did not return correct second line.")
	}
}

func TestLines2dests(t *testing.T) {
	inputc := make(chan []byte)
	retc := lines2dests(chars2lines(block2chars(inputc)), ':')
	go func() {
		inputc <-[]byte("a:2:b")
		close(inputc)
	} ()

	ret := <-retc
	if ret.source != "a" {
		t.Errorf("Error: lines2dests did not return correct source in first test.")
	}

}
