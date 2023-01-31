package fast_food

// abstract factory
type FastFoodFactory interface {
	CreateHamburger() Hamburger
	CreateChickenWing() ChickenWing
}
