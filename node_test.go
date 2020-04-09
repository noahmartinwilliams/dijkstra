package dijkstra

import "testing"
import "sync"

func TestNode(t *testing.T) {
	retc := make(chan []string)
	inputc := node(retc, "e", "e")

	go func() {
		inputc <- robot{path:[]string{"a", "e"}, pathLength:2}
		inputc <- robot{path:[]string{"s", "e"}, pathLength:1}
		close(inputc)
	} ()

	ret := <-retc
	if ret[0] != "s" {
		t.Errorf("Error: node did not return correct start name in first test.")
	}

}

func TestTreeNode(t *testing.T) {
	var dests map[string][]dest
	var wg sync.WaitGroup
	nodec := make(chan robot)

	inputc := treeNode(dests, &wg, "e", nodec)
	go func() {
		defer close(inputc)
		inputc <- robot{path:[]string{"s"}, pathLength:1}
	} ()
	node := <-nodec
	if node.path[0] != "s" {
		t.Errorf("Error: treeNode did not return proper path in first test.")
	}
}
