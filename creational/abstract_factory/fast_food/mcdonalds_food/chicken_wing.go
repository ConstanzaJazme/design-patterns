package mcdonalds_food

import (
	"fmt"
)

type ChickenWing struct{}

func (wing ChickenWing) Deliver() {
	fmt.Println("This is a chicken wing from McDonalds.")
}
