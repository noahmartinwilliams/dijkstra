package main

import "strconv"
import "os"
import "sync"

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
				panic(err.Error())
			}
			retc <- dest{dest:dest_, source:source, pathLength:pathLength}
		}
	} ()
	return retc
}

func matLine2line(wg *sync.WaitGroup, line string, separator byte, current_line int, outc chan string) {
	retc := make(chan string, 100)
	go func() {
		defer close(retc)
		num := ""
		for x := 0 ; x < len(line) ; x ++ {
			if line[x] == separator {
				retc <- num
				num = ""
			} else {
				num = num + string(line[x])
			}
		}
		retc <- num
	} ()

	go func() {
	defer wg.Done()

	start := "s"
	current_line_str := strconv.Itoa(current_line)
	if separator == 's' {
		start = "g"
	}

	x := 0
	for input := range(retc) {
		x_str := strconv.Itoa(x)
		if x_str != current_line_str {
			outc <- (start + current_line_str + string(separator) + input + string(separator) + start + x_str)
		}
		x = x + 1
	}


	} ()
}

func matLines2lines(inputc chan string, separator byte) chan string {
	retc := make(chan string, 1000)
	var wg sync.WaitGroup
	go func() {
	x := 0
	for input := range(inputc) {
		wg.Add(1)
		matLine2line(&wg, input, separator, x, retc)
		wg.Wait()
		x = x + 1
	}
	go func() {
		defer close(retc)
		wg.Wait()
	} ()
	} ()
	return retc
}
