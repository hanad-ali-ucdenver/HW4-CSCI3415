package main

import (
	"fmt"
)

func main() {
	// var strArr [3]string
	// var numArr [3]int

	// strArr[0] = "Hello World!"
	// strArr[1] = "Hello World!"
	// strArr[2] = "Hello World!"

	// numArr[0] = 1
	// numArr[1] = 2
	// numArr[2] = 3

	// fmt.Println(strArr[2])
	// fmt.Println(numArr[2])

	// var exampleSlice []int
	// exampleSlice = append(exampleSlice, 1)
	// exampleSlice = append(exampleSlice, 2)
	// exampleSlice = append(exampleSlice, 3)

	// exampleSlice2 := []int{1, 2, 3}

	// fmt.Printf("%d is the length of the Slice", len(exampleSlice))
	// fmt.Printf("%d is the capacity of the underlying array", cap(exampleSlice))

	// fmt.Printf("%d is the length of the second Slice", len(exampleSlice2))
	// fmt.Printf("%d is the capacity of the second underlying array", cap(exampleSlice2))

	// fmt.Println(exampleSlice)
	// fmt.Println(exampleSlice2)

	// other ways to make a slice exampleSlice3 := make([]int, 0, 5)
	// another way to make a slice exampleSlice4 := make([]int)

	// nameSlice := []string{"Hello", "World", "!"}

	// fmt.Println(nameSlice)

	isRaining := true

	if isRaining {
		fmt.Println("It's raining, take an umbrella")
	} else {
		fmt.Println("It's not raining. Umbrella not needed")
	}

	userName1 := "John"

	if userName1 == "John" {
		fmt.Println("You are John")
	} else {
		fmt.Println("You are not John")
	}

	if userName2 := "Jeremy"; userName2 == "John" {
		fmt.Println("You are John")
	} else {
		fmt.Println("You are not John")
	}

	a := 4
	b := 5
	sum := add(a, b)
	fmt.Println(sum)
}

func add(a int, b int) int {
	return a + b
}
