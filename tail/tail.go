package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var n int = 10

func main() {
	var file *os.File
	if len(os.Args) == 1 {
		file = os.Stdin
	} else {
		filename := os.Args[1]
		f, err := os.Open(filename)
		if err != nil {
			panic(fmt.Sprintf("failed to open file: %s, err: %v", filename, err))
		}
		file = f
	}
	buf := bufio.NewReader(file)
	out := []byte{}
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("an error ocurred while file read line: %v", err))
		}
		out = append(out, line...)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", out[len(out)-n-i])
	}
}
