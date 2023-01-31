package mcdonalds_food

import (
	"fmt"
)

type Hamburger struct{}

func (h Hamburger) Deliver() {
	fmt.Println("This is a hamburger from McDonalds.")
}
