package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func slicing_main() {

	sli := make([]int, 0, 3)

	enter := "Enter an integer"
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(enter)
		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			if text == "X" {
				break
			}

			value, err := strconv.Atoi(text)

			if err != nil {
				continue
			}

			sli = append(sli, value)

			sort.Ints(sli)
			fmt.Printf("\n%v\n", sli)
		}
	}
}
