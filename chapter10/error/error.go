package main

import (
	"errors"
	"fmt"
)

type param_error struct{}

func (object param_error) Error() string {
	return "Error, be attention!"
}

func fact(n int) (int, error) {
	if n < 0 {
		return 0, param_error{}
	}
	if n == 0 {
		return 0, errors.New("недопустимое число: должно быть положительным")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func main() {
	fact, err := fact(7)
	fmt.Println("Factorial:", fact)
	fmt.Println("Error:", err)
}
