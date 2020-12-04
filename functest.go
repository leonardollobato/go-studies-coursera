package main

func test(i int) string {
	return string(i)
}

func main() {
	fA(test)
}

func fA(fB func(int) string) {
	fB()
}

/*
func fAA() fB func(string) int{
}
*/
