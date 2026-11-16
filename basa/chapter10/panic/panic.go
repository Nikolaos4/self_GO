package main
import "fmt"

func divide_defer() {
	fmt.Println("divide_defer executed")
}

func main_defer() {
	fmt.Println("main_defer executed")
}

func divide(x, y int) int{
	defer divide_defer()
	if (y == 0) {
		panic("Division by zero!")
	}
	return x/y
}

func main() {
	defer main_defer()
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}