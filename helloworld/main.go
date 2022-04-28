package main

import (
	"fmt"

	"github.com/braejan/mentoring-golang-yz/helloworld/entities"
)

func main() {
	var counter int64 = 1
	const pi float64 = 3.141592654
	flagActive := true
	var grettings = "Hello world from string."

	fmt.Printf("%v\n", counter)
	fmt.Printf("El valor de pi es: %f\n", pi)
	fmt.Printf("El valor de flagActive es: %t\n", flagActive)
	fmt.Printf("Greet: %s\n", grettings)
	bruce := entities.NewCat("Bruce", "Solecito", 4)
	fmt.Println("pointer: ", bruce)
	fmt.Printf("value: %s\n", bruce.Name)
	brucePremium := entities.NewPremiumCat(bruce, entities.SIZE_LARGE)

	fmt.Printf("bruce premium is alive? %t\n", brucePremium.IsAlive)
	speak := bruce.Speak()
	fmt.Printf("Bruce speaking is like: %s\n", speak)
	honey := entities.NewCat("Honey", "Gray", 3)
	speak = honey.Speak()
	fmt.Printf("Honey speaking is like: %s\n", speak)

	behavior := honey.AnimalBehavior

	behavior.Speak()

}
