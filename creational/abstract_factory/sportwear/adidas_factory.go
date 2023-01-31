package sportwear

import (
	"design_patterns_in_go/creational/abstract_factory/sportwear/adidas_wear"
    "design_patterns_in_go/creational/abstract_factory/sportwear/wear"
)

type Adidas struct {
}

func (a *Adidas) makeShoe() wear.IShoe {
    return &adidas_wear.Shoe{
		Shoe: wear.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) makeShirt() wear.IShirt {
	return &adidas_wear.Shirt{
		Shirt: wear.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
