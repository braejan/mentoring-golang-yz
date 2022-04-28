package entities

// Animal is the principal struct for animal kingdom.
type Animal struct {
	IsAlive bool `json:"is_alive"`
}

type AnimalBehavior interface {
	Speak() (message string)
}

func NewAnimal(isAlive bool) (animal *Animal) {
	animal = &Animal{
		IsAlive: isAlive,
	}
	return
}
