package main

import "strconv"
import "os"

func file2blocks(stdin *os.File) chan []byte {
	retc := make(chan []byte, 100)
	buf := make([]byte, 1024)
	go func() {
		defer close(retc)
		size, err := stdin.Read(buf)
		if err != nil {
			panic(err)
		}
		if size == 0 {
			return
		}
		retc <- buf[0:size-1]
	} ()
	return retc
}

func block2chars(inputc chan []byte) chan byte {
	retc := make(chan byte, 1024)
	go func() {

	defer close(retc)
	for input := range(inputc) {
		for x := 0 ; x < len(input) ; x++ {
			retc <- input[x]
		}
	}

	} ()
	return retc
}

func chars2lines(inputc chan byte) chan string {
	retc := make(chan string, 100)
	go func() {
		defer close(retc)
		line := ""
		for input := range inputc {
			if input == '\n' {
				retc <- line
				line = ""
			} else {
				line = line + string(input)
			}
		}
		if line != "" {
			retc <-line
		}

	} ()
	return retc
}

func lines2dests(inputc chan string, separator byte) chan dest {
	retc := make(chan dest, 100)
	go func() {
		defer close(retc)
		for input := range(inputc) {
			source := ""
			dest_ := ""
			pathLengthStr := ""
			x := 0
			for ; input[x] != separator && x < len(input) ; x++ {
				source = source + string(input[x])
			}

			for x = x + 1 ; input[x] != separator && x < len(input) ; x++ {
				pathLengthStr = pathLengthStr + string(input[x])
			}

			for x = x + 1 ; x < len(input) ; x++ {
				dest_ = dest_ + string(input[x])
			}

			pathLength, err := strconv.ParseFloat(pathLengthStr, 64)
			if err != nil {
				continue
			}
			retc <- dest{dest:dest_, source:source, pathLength:pathLength}
		}
	} ()
	return retc
}
