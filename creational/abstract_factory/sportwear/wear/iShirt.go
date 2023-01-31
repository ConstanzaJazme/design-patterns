package wear

import "fmt"

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
    PrintDetails()
}

type Shirt struct {
	Logo string
	Size int
}

func (s *Shirt) setLogo(logo string) {
	s.Logo = logo
}

func (s *Shirt) getLogo() string {
	return s.Logo
}

func (s *Shirt) setSize(size int) {
	s.Size = size
}

func (s *Shirt) getSize() int {
	return s.Size
}

func (s Shirt) PrintDetails() {
	fmt.Printf("[Shirt] Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("[Shirt] Size: %d", s.getSize())
	fmt.Println()
}
