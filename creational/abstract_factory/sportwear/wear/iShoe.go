package wear

import "fmt"

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
    PrintDetails()
}

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) setLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) getLogo() string {
	return s.Logo
}

func (s *Shoe) setSize(size int) {
	s.Size = size
}

func (s *Shoe) getSize() int {
	return s.Size
}

func (s Shoe) PrintDetails() {
    fmt.Printf("[Shoe] Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("[Shoe] Size: %d", s.getSize())
    fmt.Println()
}
