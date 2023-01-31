package sportwear

import (
	"design_patterns_in_go/creational/abstract_factory/sportwear/wear"
	"fmt"
)

type ISportsFactory interface {
	makeShoe() wear.IShoe
	makeShirt() wear.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
