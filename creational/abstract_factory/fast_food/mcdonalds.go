package fast_food

import "design_patterns_in_go/creational/abstract_factory/fast_food/mcdonalds_food"

// concrete factory two: McDonalds
type Mcdonalds struct{}

func (factory Mcdonalds) CreateHamburger() Hamburger {
    return new(mcdonalds_food.Hamburger)
}
func (factory Mcdonalds) CreateChickenWing() ChickenWing {
    return new(mcdonalds_food.ChickenWing)
}
