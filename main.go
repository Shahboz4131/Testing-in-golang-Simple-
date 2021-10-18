package main

import "fmt"

func Addition(x, y int) int {
	return x + y
}

func Subtraction(x, y int) int {
	return x - y
}

func Multiplication(x, y int) int {
	return x * y
}

func Division(x, y int) (result int, message string, err bool) {
	if y == 0 {
		message = "Error y can not be 0"
	}else {
		err = true
		result = x / y
	}
	return result, message, err
}

func main () {

	a := Addition(2, 5)
	fmt.Println(a)

	b := Subtraction(3, 1)
	fmt.Println(b)

	c := Multiplication(4, 6)
	fmt.Println(c)

	result, message, err := Division(4, 0)
	if err {
		fmt.Println(result)
	}else {
		fmt.Println(message)
	}
}