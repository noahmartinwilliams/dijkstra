package main

import "testing"
import "sync"

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

func TestMatLine2line(t *testing.T) {
	var wg sync.WaitGroup
	outc := make(chan string)
	wg.Add(1)
	matLine2line(&wg, "0:1:1:3", ':', 0, outc)
	out := <-outc
	if out != "s0:1:s1" {
		t.Errorf("Error: matLine2line failed to return correct first output. Got: \"" + out + "\"")
	}
}

func TestMatLines2lines(t *testing.T) {
	inputc := make(chan string)
	retc := matLines2lines(inputc, ':')
	go func() {
		inputc <- "0:1:2"
		inputc <- "1:0:2"
	} ()

	ret0 := <-retc
	if ret0 != "s0:1:s1" {
		t.Errorf("Error: matLine2lines failed to return correct first output. Got: \"" + ret0 + "\"")
	}

	ret1 := <-retc
	if ret1 != "s0:2:s2" {
		t.Errorf("Error: matLine2lines failed to return correct second output. Got: \"" + ret1 + "\"")
	}

	ret2 := <-retc
	if ret2 != "s1:1:s0" {
		t.Errorf("Error: matLine2lines failed to return correct third output. Got: \"" + ret2 + "\"")
	}

	ret3 := <-retc
	if ret3 != "s1:2:s2" {
		t.Errorf("Error: matLine2lines failed to return correct fourth output. Got: \"" + ret3 + "\"")
	}
}
