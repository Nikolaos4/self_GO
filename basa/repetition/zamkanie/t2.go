package main
import "fmt"

func multiplier(factor int) func(int) int{
	return func(x int) int {return x * factor}
}

func main() {
	f := multiplier(4)(6)
	fmt.Println(f)
}