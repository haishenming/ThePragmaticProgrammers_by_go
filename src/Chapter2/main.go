package main

import (
	"fmt"
	oper "Chapter2/Operation"
	retire "Chapter2/Retire"
)

func T1() {
	// say hello
	var name string

	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)

	fmt.Printf("Hello %s, nice to see you!\n", name)
}

func T2() {
	// count str
	var str string

	fmt.Println("What is the input string?")
	fmt.Scanf("%s", &str)

	fmt.Printf("%s has %d characters.\n", str, len(str))
}

func T3() {
	var quote,author string

	fmt.Println("What is the quote?")
	fmt.Scanf("%s", &quote)
	fmt.Println("Who said it?")
	fmt.Scanf("%s", &author)

	fmt.Printf("%s said: \"%s\"", author, quote)
}

func T4() {
	var noun,verb, adj, adv string

	fmt.Println("Enter a noun:")
	fmt.Scanf("%s", &noun)

	fmt.Println("Enter a verb:")
	fmt.Scanf("%s", &verb)

	fmt.Println("Enter a adj:")
	fmt.Scanf("%s", &adj)

	fmt.Println("Enter a adv:")
	fmt.Scanf("%s", &adv)

	fmt.Printf("Do you %s your %s %s %s? That's hiarious!\n", verb, adj, noun, adv)
}

func T5()  {

	var n1, n2 int

	fmt.Println("Waht is the first number?")
	fmt.Scanf("%d", &n1)
	fmt.Println("What is the second number?")
	fmt.Scanf("%d", &n2)

	add, sub, mul, div := oper.Operation(n1, n2)
	fmt.Printf("%d + %d = %.2f\n" +
						"%d - %d = %.2f\n" +
						"%d * %d = %.2f\n" +
						"%d / %d = %.2f\n", n1, n2, add, n1, n2, sub, n1, n2, mul, n1, n2, div)

}

func T6() {

	var age, retireAge int
	fmt.Println("What is you current age?")
	fmt.Scanf("%d", &age)

	fmt.Println("What age would you like retire?")
	fmt.Scanf("%d", &retireAge)

	hasYear, retireYear, currentYear := retire.OperationRetire(age, retireAge)

	fmt.Printf("You have %d years left untill you can retire.\n", hasYear)
	fmt.Printf("It's %d, so you can retire in %d", currentYear, retireYear)

}

func main() {
	var T int
	fmt.Println("Which T would you want? (You olny need to enter the number!)")
	fmt.Scanf("%d", &T)

	switch T {
	case 1:
		fmt.Println("Start T1")
		T1()
	case 2:
		fmt.Println("Start T2")
		T2()
	case 3:
		fmt.Println("Start T3")
		T3()
	case 4:
		fmt.Println("Start T4")
		T4()
	case 5:
		fmt.Println("Start T5")
		T5()
	case 6:
		fmt.Println("Start T6")
		T6()
	default:
		fmt.Printf("Not find your T%d\n", T)
	}

}
