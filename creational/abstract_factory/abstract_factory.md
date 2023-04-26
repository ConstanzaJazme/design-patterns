# Abstract Factory

## Fast Food example

```mermaid
classDiagram
`FastFoodFactory (abstract factory)` <|-- `KFC (concrete factory)`
`FastFoodFactory (abstract factory)` <|-- `McDonalds (concrete factory)`
`KFC (concrete factory)` 
`McDonalds (concrete factory)` 

Class03 *-- Class04
Class05 o-- Class06
Class07 .. Class08
Class09 --> C2 : Where am i?
Class09 --* C3
Class09 --|> Class07
Class07 : equals()
Class07 : Object[] elementData
Class08 <--> C2: Cool label

class `FastFoodFactory (abstract factory)` {
        ...
        +CreateHamburger()
        +CreateChickenWing()
    }

class `KFC (concrete factory)` {
        ...
        +CreateHamburger()
        +CreateChickenWing()
    }

class `McDonalds (concrete factory)` {
        ...
        +CreateHamburger()
        +CreateChickenWing()
    }
