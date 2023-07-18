package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

// cow
type Cow struct {
	food, locomotion, noise string
}

func (c Cow) Eat()   { fmt.Println(c.food) }
func (c Cow) Move()  { fmt.Println(c.locomotion) }
func (c Cow) Speak() { fmt.Println(c.noise) }

// bird
type Bird struct {
	food, locomotion, noise string
}

func (b Bird) Eat()   { fmt.Println(b.food) }
func (b Bird) Move()  { fmt.Println(b.locomotion) }
func (b Bird) Speak() { fmt.Println(b.noise) }

// snake
type Snake struct {
	food, locomotion, noise string
}

func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func main() {
	animals := map[string]Animal{}
	for {
		var request string
		var name string
		var animalOrInfoName string
		fmt.Println("> ")
		fmt.Scanln(&request, &name, &animalOrInfoName)
		if request == "newanimal" {
			switch animalOrInfoName {
			case "cow":
				animals[name] = Cow{food: "grass", locomotion: "walk", noise: "moo"}
			case "bird":
				animals[name] = Bird{food: "worms", locomotion: "fly", noise: "peep"}
			case "snake":
				animals[name] = Snake{food: "mice", locomotion: "slither", noise: "hss"}
			}
			fmt.Println("Created it!")
		} else {
			switch animalOrInfoName {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			}
		}
	}
}

func funcA() func(int) string {
	return func(int) string { return "" }
}
