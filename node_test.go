package main

import "testing"
//import "sync"

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

/*func TestTreeNode(t *testing.T) {
	dests := make( map[string][]dest )
	nodePool := make(map[string]chan robot)
	var wg sync.WaitGroup
	wg.Add(1)
	nodePool["e"] = make(chan robot)

	inputc := treeNode(dests, &wg, "e", nodePool)
	go func() {
		defer close(inputc)
		inputc <- robot{path:[]string{"s"}, pathLength:1}
	} ()
	node := <-nodePool["e"]
	if node.path[0] != "s" {
		t.Errorf("Error: treeNode did not return proper path in first test.")
	}

	var wg2 sync.WaitGroup
	wg2.Add(1)
	dests["a"]=[]dest{dest{dest:"e", pathLength:1}}
	nodePool["e"] = make(chan robot)
	inputc = treeNode(dests, &wg2, "a", nodePool)

	go func() {
		defer close(inputc)
		inputc <- robot{path:[]string{"s"}, pathLength:1}
	} ()

	node = <-nodePool["e"]
	if node.path[2] != "e" {
		t.Errorf("Error: treeNode did not launch subTree.")
	}
} */
