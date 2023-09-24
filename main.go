package main

import "fmt"

func addFirstTwoAndStore(arr *[6]int) {
	arr[0] = arr[0] + arr[1]
}

// Define a function type for our map
type MenuFunction func()

func option1() {
	fmt.Println("You selected option 1!")
}

func option2() {
	fmt.Println("You selected option 2!")
}

func option3() {
	fmt.Println("You selected option 3!")
}

func main() {
	myArray := [6]int{1, 2, 3, 4, 5, 6}
	options := [3]func(){option1, option2, option3}

	var userInput int

	fmt.Println("Before:", myArray)

	fmt.Println("Player 1: ")

	fmt.Println("Select an option (1-3):")
	fmt.Scanln(&userInput)

	// Use the modulus operator to make sure userInput is a valid index
	// The "+ len(options)" ensures the userInput is not negative
	userInput = (userInput + len(options)) % len(options)

	options[userInput]()

	//checkwinCondition()

	fmt.Println("After:", myArray)

	fmt.Println("Player 2: ")

	//checkWinCondition()

}
