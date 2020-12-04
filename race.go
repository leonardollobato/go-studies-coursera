package main

var balance int

/*
race condition 1: if we use go routines here
a race condition happens because the final value
of balance depends on the proper order in which its executed.
*/
func transaction(amount int) {
	balance = balance + amount
}

/*
race condition 2: if we use go routines here
a race condition happens because you cannot
determine the real order of the message since
you don't know when go routine is gonna be executed
*/
func room(message string, chat *[]string) {
	*chat = append(*chat, message)
}

/*
func main() {

	// Race Condition 1
	fmt.Println(balance)
	go transaction(100)
	fmt.Println(balance)
	go transaction(-50)
	fmt.Println(balance)

	// Race condition 2
	chat := make([]string, 0)

	go room("john: hello anne", &chat)
	go room("anne: hello john", &chat)
	go room("john: how are you anne?", &chat)
	go room("anne: I am great john, how about you?", &chat)

	for _, message := range chat {
		fmt.Println(message)
	}

}
*/
