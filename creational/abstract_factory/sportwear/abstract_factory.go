package sportwear

func Start() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

    nikeShoe.PrintDetails()
    nikeShirt.PrintDetails()

    adidasShoe.PrintDetails()
    adidasShirt.PrintDetails()
}
