package main
import "fmt"

func calculate(x, y int) (int, int, int){
	return (x+y), (x-y), (x*y);
}

func main() {
	var x, y int;
	fmt.Scan(&x);
	fmt.Scan(&y);
	sum, raz, mult := calculate(x, y);
	fmt.Println(sum, raz, mult);
}