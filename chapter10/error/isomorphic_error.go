package main
import "fmt"

func factorial(n int) (int, interface{}){
	if (n < 0) {
		return 0, "Некорректное число"
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func main() {
	n := -7
	fac, err := factorial(n)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(fac)
}