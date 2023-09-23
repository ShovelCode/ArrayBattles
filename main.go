package main

import "fmt"

func addFirstTwoAndStore(arr *[5]int) {
    arr[0] = arr[0] + arr[1]
}

func main() {
    myArray := [5]int{1, 2, 3, 4, 5}

    fmt.Println("Before:", myArray)  // Outputs: [1 2 3 4 5]

    addFirstTwoAndStore(&myArray)

    fmt.Println("After:", myArray)   // Outputs: [3 2 3 4 5]
}
