package main

import (
	"fmt"
	Area "Chapter3/Area"
)
func T7()  {
	var length, width int

	fmt.Println("What is the length of the room in feet?")
	fmt.Scanf("%d", &length)

	fmt.Println("What is the width of the room in feet?")
	fmt.Scanf("%d", &width)

	feetArea, meterArea := Area.Area(length, width)

	fmt.Printf("%d square feet\n", feetArea)
	fmt.Printf("%.2f square meters", meterArea)
}

func main() {
	var T int
	fmt.Println("Which T would you want? (You olny need to enter the number!)")
	fmt.Scanf("%d", &T)

	switch T {
	case 7:
		fmt.Println("Start T1")
		T7()

	default:
		fmt.Printf("Not find your T%d\n", T)
	}
}
