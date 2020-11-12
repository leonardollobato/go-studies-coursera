package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func marshalling_main() {

	p1 := Person{firstName: "Leonardo", lastName: "Lobato", age: 33}
	var p2 Person

	barr, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("Something went wrong")
	}

	fmt.Printf("data: %v", barr)

	err = json.Unmarshal(barr, &p2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nfirst name: %s", p2.firstName)
}

type Person struct {
	firstName string
	lastName  string
	age       int
}
