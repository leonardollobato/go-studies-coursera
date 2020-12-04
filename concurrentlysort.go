package main

/*
const coef = 4

func main() {
	fmt.Println("Type at least 4 integers values: ")
	var chunks [][]int
	var finalarr []int

	inputs, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputlen := len(inputs)
	mod := inputlen % coef
	x := inputlen / coef

	start := 0
	end := x

	c := make(chan []int, inputlen)

	for i := 0; i < coef; i++ {
		chunks = append(chunks, inputs[start:end])
		start = end
		if mod > 0 {
			end = end + x + 1
			mod--
		} else {
			end = end + x
		}
	}

	for i := 0; i < len(chunks); i++ {
		go sort(chunks[i], c)
		sortedarr := <-c
		for _, v := range sortedarr {
			finalarr = append(finalarr, v)
		}
	}

	go sort(finalarr, c)
	fmt.Println("Sorted Array:", <-c)
}

func sort(items []int, c chan []int) {
	size := len(items)
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

	c <- items
}

func Swap(items []int, index int) {
	temp := items[index]
	items[index] = items[index-1]
	items[index-1] = temp
}

func readInput() ([]int, error) {
	var i []int
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()

	inputs := strings.Split(text, " ")
	if len(inputs) < 4 {
		err = errors.New("Four values are required")
		return nil, err
	}

	for _, value := range inputs {
		input, err := strconv.Atoi(value)

		if err != nil {

			err = errors.New("All values need to be integer")
			return nil, err
		}

		i = append(i, input)
	}

	return i, nil
}
*/
