package main

import "fmt"

func main() {
	fmt.Println("Hello new language")

	var number1 int
	fmt.Println(number1)

	var number2 = 10
	fmt.Println(number2)

	number3 := 30
	fmt.Println(number3)

	number := 45
	fmt.Println("initial value", number)

	number = 95
	fmt.Println("the changed value", number)

	var name, age = "Evelina", 21
	fmt.Println("Name", name, "age", age)

	name1, age1 := "Andrew", 20
	fmt.Println(name1, age1)

	const pi = 3.15
	fmt.Println(pi)

	const comName = "M&M"
	fmt.Println(comName)

	currentAge := 21
	fmt.Printf("Age - %d years old", currentAge)

	var enterName string

	fmt.Println("\nEnter ur name")
	fmt.Scan(&enterName)

	fmt.Printf("U enter - %s\n", enterName)
}
