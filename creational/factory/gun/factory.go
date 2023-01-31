package main

import (
    "fmt"
    . "gun/types"
)

func NewGun(gunType string) (IGun, error) {
    if gunType == "ak47" {
        return NewAk47(), nil
    }
    if gunType == "musket" {
        return NewMusket(), nil
    }
    return nil, fmt.Errorf("wrong gun type passed")
}

func GetGun(g IGun) {
    fmt.Printf("Gun: %s", g.GetName())
    fmt.Println()
    fmt.Printf("Power: %d", g.GetPower())
    fmt.Println()
}