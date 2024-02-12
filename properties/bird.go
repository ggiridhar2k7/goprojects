package properties

import "fmt"

var birdProperties = map[string]string{
	"food":  "worms",
	"loco":  "fly",
	"speak": "peep",
}

type Bird struct {
	name string
}

func (b *Bird) Eat() {
	fmt.Println("Bird eats", birdProperties["food"])
}

func (b *Bird) Move() {
	fmt.Println("Bird moves by", birdProperties["loco"])
}

func (b *Bird) Speak() {
	fmt.Println("Bird speaks", birdProperties["speak"])
}

func (b *Bird) GetName() string {
	return b.name
}
