package properies

import "fmt"

type AnimalProperty int

const (
	Food AnimalProperty = iota
	Movement
	Sound
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

var userAnimalDB map[Animal]string

func GetProperty(a Animal, prop string) {
	switch prop {
	case "food":
		a.Eat()
	case "loco":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Unsupported property of the Animal, pls enter one among ('food', 'loco', 'speak')")
	}
}

func UpdateAnimalName(a Animal, animalType string) {
	userAnimalDB[a.GetName()] = animalType
}

func GetAnimalType(name string) string {
	result, _ := userAnimalDB[name]
	return result
}
