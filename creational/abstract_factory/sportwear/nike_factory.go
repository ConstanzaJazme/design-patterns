package sportwear

import (
    "design_patterns_in_go/creational/abstract_factory/sportwear/nike_wear"
    "design_patterns_in_go/creational/abstract_factory/sportwear/wear"
)

type Nike struct {
}

func (n *Nike) makeShoe() wear.IShoe {
	return &nike_wear.Shoe{
		Shoe: wear.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) makeShirt() wear.IShirt {
	return &nike_wear.Shirt{
		Shirt: wear.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
