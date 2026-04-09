package main
import "fmt"

func main() {
	fmt.Println(add([]int{1,4,2,5}...));
}

func add(numbers ...int) int {
	sum := 0;
	for _, value :=  range numbers {
		sum += value;
	}
	return sum
}