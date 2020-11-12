package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func findian_main() {

	var stdin *bufio.Reader
	var line []rune
	var c rune
	var err error
	stdin = bufio.NewReader(os.Stdin)

	for {
		c, _, err = stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
			os.Exit(1)
		}
		line = append(line, c)
	}

	input := strings.ToLower(strings.TrimSpace(string(line)))

	first := input[:1]
	last := input[len(input)-1:]
	between := input[1 : len(input)-1]

	if last == "n" &&
		first == "i" &&
		strings.Contains(between, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
