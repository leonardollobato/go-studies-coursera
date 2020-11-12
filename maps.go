package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Info struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	info := make(map[string]string)

	messages := map[string]string{"name": "Enter your name:", "address": "Enter your address:"}
	scanner := bufio.NewScanner(os.Stdin)

	for key, element := range messages {
		fmt.Println(element)

		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			info[key] = text
		}
	}

	//i := Info{Name: info["name"], Address: info["address"]}

	json, _ := json.Marshal(info)
	fmt.Println(string(json))
}
