package main

import "os"
import "fmt"
import "github.com/pborman/getopt"

func usage() {
	fmt.Println("dijkstra - run dijkstra's algorithm.")
	fmt.Println("The graph should be supplied in the format of lines that each contain one path from a start node to an end node.")
	fmt.Println("The lines should be of the format \"start:distance:end\" (assuming that the colon character is the separator.")
	fmt.Println("Use -s to specify the starting node, -e to specify the ending node, and -d to specify what the delimiter should be between elements of a line.")
	fmt.Println("Use -m to specify a matrix as input with rows terminated by the newline character, and columns terminated by separator character nodes will be autolabeled with 's' as the start character unless 's' is the delimiter in which case 'g' will be used.")
	fmt.Println("Use -h for help.")
}

func main() {
	separator := getopt.String('d', ":", "Specify what the delimiter character between parts of a line is.")
	start := getopt.String('s', "a", "Specify where to start.")
	end := getopt.String('e', "e", "Specify where to end.")
	help := getopt.Bool('h', "Print help message.")
	use_mat := getopt.Bool('m', "Take input as matrix")

	getopt.SetUsage(usage)
	getopt.Parse()

	if *help {
		usage()
		return
	}

	c0  := chars2lines(block2chars(file2blocks(os.Stdin)))
	sep_char := (*separator)[0]
	if *use_mat {
		c := lines2dests(matLines2lines(c0, sep_char), sep_char)
		ret := launchTreeNodes(*end, *start, c)
		for x := 0 ; x < len(ret) ; x++ {
			fmt.Print(ret[x])
			if x != len(ret) - 1 {
				fmt.Print(string((*separator)[0]))
			}
		}
		//TODO: figure out why it fails to print out the last node in matrix mode.
		fmt.Print(string((*separator)[0]))
		fmt.Print(*end)
	} else {
		c := lines2dests(c0, sep_char)
		ret := launchTreeNodes(*end, *start, c)
		for x := 0 ; x < len(ret) ; x++ {
			fmt.Print(ret[x])
			if x != len(ret) - 1 {
				fmt.Print(string((*separator)[0]))
			}
		}
	}
}
