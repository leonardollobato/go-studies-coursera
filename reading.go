package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const filelocation = "./info.txt"

type Persona struct {
	fname string
	lname string
}

func reading_main() {
	var infos []Persona

	f, err := os.Open(filelocation)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fname := line[0]
		lname := line[1]
		if len(fname) > 20 || len(lname) > 20 {
			log.Fatal("maximum string size is 20")
		}
		infos = append(infos, Persona{fname, lname})
	}

	for _, v := range infos {
		fmt.Printf("%s %s\n", v.fname, v.lname)
	}

}
