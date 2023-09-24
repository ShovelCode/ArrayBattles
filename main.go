package main

import "fmt"

func addFirstTwoAndStore(arr *[6]int) {
	arr[0] = arr[0] + arr[1]
}

// Define a function type for our map
type MenuFunction func()

func option1(arr *[6]int) {
	fmt.Println("You selected option 1!")
}

func option2(arr *[6]int) {
	fmt.Println("You selected option 2!")
}

func option3(arr *[6]int) {
	fmt.Println("You selected option 3!")
}

func option4(arr *[6]int) {
	arr[0] = arr[1] * arr[2];

}

func mutateArray(arr *[3]int) {
    for i := range arr {
        arr[i] *= 2
    }
}
func main() {
	myArray := [6]int{1, 2, 3, 4, 5, 6}
	options := [4]func(arr *[6]int){option1, option2, option3, option4}

	var userInput int

	fmt.Println("Before:", myArray)

	fmt.Println("Player 1: ")

	fmt.Println("Select an option (1-4):")
	fmt.Scanln(&userInput)

	
	userInput = (userInput + len(options)) - len(options)

	options[userInput](&myArray)

	//checkwinCondition()

	fmt.Println("After:", myArray)

	fmt.Println("Player 2: ")

	//checkWinCondition()

}
