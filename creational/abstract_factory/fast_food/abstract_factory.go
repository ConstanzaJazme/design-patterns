package fast_food

import (
	"flag"
	"fmt"
	"os"
)

func Start() {
	prefer := GetPreferredFastFood()

	var factory FastFoodFactory
	switch prefer {
	case "KFC":
		factory = new(Kfc)
	case "McDonalds":
		factory = new(Mcdonalds)
	default:
		fmt.Printf("%s not supported yet.\n", prefer)
		os.Exit(1)
	}

	hamburger := factory.CreateHamburger()
	chickenWing := factory.CreateChickenWing()

	hamburger.Deliver()
	chickenWing.Deliver()
}

func GetPreferredFastFood() string {
	var prefer string
    flag.StringVar(&prefer, "prefer", "McDonalds", "preferred fast food type")
	flag.Parse()

//	if len(os.Args) != 3 {
//		flag.Usage()
//		os.Exit(1)
//	}

	return prefer
}