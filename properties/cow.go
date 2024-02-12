package properties

import "fmt"

var cowProperties = map[string]string{
	"food":  "grass",
	"loco":  "walk",
	"speak": "moo",
}

type Cow struct {
	name string
}

func (c *Cow) Eat() {
	fmt.Println("Cow eats", cowProperties["food"])
}

func (c *Cow) Move() {
	fmt.Println("Cow moves by", cowProperties["loco"])
}

func (c *Cow) Speak() {
	fmt.Println("Cow speaks", cowProperties["speak"])
}
