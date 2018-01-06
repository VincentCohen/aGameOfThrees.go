package main

import (
	"fmt"
)

func main() {
	// Ask the question
	str := "Enter a number: "
	fmt.Print(str)

	var num int
	fmt.Scanf("%d", &num)

	calc(num)
}

func calc(number int) {
	if number == 1 {
		fmt.Println(number, "END")
		return
	}

	if isDividable(number) {

		fmt.Println(number, "0")

		number = divide(number)

		calc(number)
	} else {
		if isDividable(number - 1) {
			fmt.Println(number, "-1")

			number = divide(minus(number))

			calc(number)
		} else {
			fmt.Println(number, "+1")

			number = divide(add(number))

			calc(number)
		}
	}
}

func isDividable(number int) bool {
	return number%3 == 0
}

func divide(number int) int {
	return number / 3
}

func minus(number int) int {
	return number - 1
}

func add(number int) int {
	return number + 1
}
