package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	const Truth = true
	fmt.Println("Go rules?", Truth)

	const Pi = 3.14
	const World = 2
	fmt.Println("Hello", World)
	fmt.Println("Hien so", Pi)
}
