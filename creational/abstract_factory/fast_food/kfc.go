package fast_food

import "design_patterns_in_go/creational/abstract_factory/fast_food/kfc_food"

// concrete factory one: KFC
type Kfc struct{}

func (factory Kfc) CreateHamburger() Hamburger {
	return new(kfc_food.Hamburger)
}
func (factory Kfc) CreateChickenWing() ChickenWing {
	return new(kfc_food.ChickenWing)
}
