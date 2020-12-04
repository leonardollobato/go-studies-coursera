package main

/*
type Animal interface {
	Speak()
	Move()
	Eat()
}

type ZooAnimal struct {
	eat, move, sound string
}

func (a ZooAnimal) Eat() {
	fmt.Println(a.eat)
}

func (a ZooAnimal) Move() {
	fmt.Println(a.move)
}

func (a ZooAnimal) Speak() {
	fmt.Println(a.move)
}

func main() {
	animals := make(map[string]Animal)
	defaultanimals := map[string]Animal{
		"cow":   ZooAnimal{eat: "grass", move: "walk", sound: "moo"},
		"bird":  ZooAnimal{eat: "worms", move: "fly", sound: "peep"},
		"snake": ZooAnimal{eat: "mice", move: "slither", sound: "hsss"},
	}

	for {
		askForInput()
		var requesttype, name, info string
		fmt.Scan(&requesttype, &name, &info)

		switch requesttype {
		case "newanimal":
			newanimal, ok := defaultanimals[info]
			if !ok {
				fmt.Println("this animal type is not recognized")
				continue
			}
			animals[name] = newanimal
			fmt.Println("Created it!")
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Println("Unknow Animal:", name)
				continue
			}
			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknow Info: ", info)
			}
		default:
			fmt.Println("this command is not recognized")
		}
	}

}

func askForInput() {
	fmt.Print(">")
}*/
