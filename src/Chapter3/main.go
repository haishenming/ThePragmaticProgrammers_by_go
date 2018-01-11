package main

import (
	"fmt"
	Area "Chapter3/Area"
	Pizza "Chapter3/Pizza"
)
func T7()  {
	var length, width int

	fmt.Println("What is the length of the room in feet?")
	fmt.Scanf("%d", &length)

	fmt.Println("What is the width of the room in feet?")
	fmt.Scanf("%d", &width)

	feetArea, meterArea := Area.Area(length, width)

	fmt.Printf("%d square feet\n", feetArea)
	fmt.Printf("%.2f square meters\n", meterArea)
}

func T8() {
	var pizzas, people, pieces int

	fmt.Println("How many people?")
	fmt.Scanf("%d", &people)

	fmt.Println("How many pizzas do you have?")
	fmt.Scanf("%d", &pizzas)

	fmt.Println("Each pizza divided into several pieces?")
	fmt.Scanf("%d", &pieces)

	peoplePieces, leftover := Pizza.DividedPizzas(pizzas, people, pieces)

	fmt.Printf("%d people with %d pizzas.\n", people, pizzas)
	fmt.Printf("Each people gets %d pieces of pizzas.\n", peoplePieces)
	fmt.Printf("There are %d leftover pieces.\n", leftover)

}

func main() {
	var T int
	fmt.Println("Which T would you want? (You olny need to enter the number!)")
	fmt.Scanf("%d", &T)

	switch T {
	case 7:
		fmt.Println("Start T7")
		T7()
	case 8:
		fmt.Println("Start T8")
		T8()

	default:
		fmt.Printf("Not find your T%d\n", T)
	}
}
