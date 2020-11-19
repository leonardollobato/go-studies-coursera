package sorting

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const numberofinputs = 10
const prompt = "Enter a number (%d/%d)"

func bubblesort_main() {
	var counter = 0
	inputlist := make([]int, 0, numberofinputs)
	scanner := bufio.NewScanner(os.Stdin)

	for counter < numberofinputs {
		AskForInput(&counter, numberofinputs)
		input, err := strconv.Atoi(ReadInput(scanner))

		if err != nil {
			log.Fatal("Error when trying to convert read input to number")
		}

		inputlist = append(inputlist, input)
		counter += 1
	}

	BubbleSort(inputlist, numberofinputs)

	fmt.Printf("%v", inputlist)
}

func BubbleSort(items []int, size int) {
	isSwapped := true

	for isSwapped {
		isSwapped = false

		for i := 1; i < size; i++ {
			if items[i-1] > items[i] {
				Swap(items, i)
				isSwapped = true
			}
		}
	}
}

func Swap(items []int, index int) {
	temp := items[index]
	items[index] = items[index-1]
	items[index-1] = temp
}

func ReadInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func AskForInput(counter *int, sizeinput int) {
	fmt.Println(fmt.Sprintf(prompt, *counter+1, sizeinput))
}
