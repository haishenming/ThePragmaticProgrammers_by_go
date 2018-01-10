package main

import "fmt"

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

func main() {
	T1()
	T2()
}
