package entities

import "fmt"

type Cat struct {
	*Animal
	Name  string `json:"name"`
	Color string `json:"color"`
	Age   uint8  `json:"age"`
	AnimalBehavior
}

func (cat *Cat) Speak() (message string) {
	message = fmt.Sprintf("Hi mi name is %s and I am %d years old.", cat.Name, cat.Age)
	return
}

func NewCat(name string, color string, age uint8) (cat *Cat) {
	newAnimalObject := NewAnimal(true)
	cat = &Cat{
		Animal: newAnimalObject,
		Name:   name,
		Color:  color,
		Age:    age,
	}
	cat.AnimalBehavior = cat
	return
}
