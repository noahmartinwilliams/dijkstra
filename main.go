package main

import "os"
import "fmt"
import "github.com/pborman/getopt"

func usage() {
	fmt.Println("dijkstra - run dijkstra's algorithm.")
	fmt.Println("The graph should be supplied in the format of lines that each contain one path from a start node to an end node.")
	fmt.Println("The lines should be of the format \"start:distance:end\" (assuming that the colon character is the separator.")
	fmt.Println("Use -s to specify the starting node, -e to specify the ending node, and -d to specify what the delimiter should be between elements of a line.")
	fmt.Println("Use -h for help.")
}

func main() {
	separator := getopt.String(100, ":", "Specify what the delimiter character between parts of a line is.")
	start := getopt.String(115, "a", "Specify where to start.")
	end := getopt.String(101, "e", "Specify where to end.")
	help := getopt.Bool(104, "Print help message.")
	getopt.SetUsage(usage)
	getopt.Parse()

	if *help {
		usage()
		return
	}

	c := lines2dests(chars2lines(block2chars(file2blocks(os.Stdin))), (*separator)[0])
	ret := launchTreeNodes(*end, *start, c)
	for x := 0 ; x < len(ret) ; x++ {
		fmt.Print(ret[x])
		if x != len(ret) - 1 {
			fmt.Print(string((*separator)[0]))
		}
	}
}
