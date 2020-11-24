package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const basemessage = "Please inform %s (float64): "

func func_main() {
	//s =½ a t2 + vot + so

	messages := [3]string{"acceleration", "initial velocity", "initial displacement"}
	scanner := bufio.NewScanner(os.Stdin)
	inputlist := make([]float64, 0, len(messages))

	for _, v := range messages {
		AskForInput(v)
		input, err := strconv.ParseFloat(ReadInput(scanner), 64)

		if err != nil {
			log.Fatal("Error when trying to convert read input to number")
		}

		inputlist = append(inputlist, input)
	}

	if len(inputlist) == 3 {

		s := GetDisplaceFn(inputlist[0], inputlist[1], inputlist[2])
		fmt.Println(fmt.Sprintf("Displacement after %d seconds: %f", 3, s(3)))
		fmt.Println(fmt.Sprintf("Displacement after %d seconds: %f", 5, s(5)))

	} else {
		log.Fatal("Number of arguments are wrong")
	}
}

func GetDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return .5*a*(math.Pow(t, 2)) + v0*t + s0
	}
}

func AskForInput(message string) {
	fmt.Println(message)
}

func ReadInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
