package main

import "fmt"

func main() {
	var name = "Abdullah"
	var age int = 21
	isStudent := true
	fmt.Println(name, age, isStudent)

	var num1, num2 = 10, 20
	fmt.Println(num1, num2)

	var firstName, lastName string = "Abdullah", "Zahid"
	fmt.Println(firstName, lastName)

	skill1, skill2 := "Java", "Go"
	fmt.Println(skill1, skill2)

	var school string
	school = "UET"
	fmt.Println(school)

	const pi = 3.14
	fmt.Println(pi)
}
