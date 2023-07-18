package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat()   { fmt.Println(a.food) }
func (a Animal) Move()  { fmt.Println(a.locomotion) }
func (a Animal) Speak() { fmt.Println(a.noise) }

func main() {
	// animals
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}

	// infinite loop
	for {
		// prompt user for input
		var animal string
		var request string
		fmt.Println("> ")
		fmt.Scanln(&animal, &request)
		// match on animal and request
		switch animal {
		case "cow":
			switch request {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "snake":
			switch request {
			case "eat":
				snake.Eat()
			case "move":
				snake.Eat()
			case "speak":
				snake.Eat()
			}
		case "bird":
			switch request {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		}
	}
}
