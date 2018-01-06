package main

import (
	"fmt"
)

func main() {
	// Ask the question
	str := "Enter a number to divide: "
	fmt.Print(str)

	var num int
	fmt.Scanf("%d", &num)

	fmt.Println("You selected:", num)
}
