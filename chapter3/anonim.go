package main
import "fmt"

func add(x int, y int, operation func(int, int) int) int{
	return x*y + operation(x, y);
}

func get_func(n int) func(int, int) int{
	switch(n) {
		case 1: 
			return func(x int, y int) int{return x+y};	
		case 2:
			return func(x int, y int) int{return x*y}
		default: 
			return func(x, y int) int{return x/y}
	}
}

func main() {
	f := add(2, 3, func(x int, y int) int {return x + y});	
	fmt.Println(f);

	var n int;
	fmt.Scanln(&n);

	f2 := get_func(n);
	fmt.Print("\n",f2(3,4));
}