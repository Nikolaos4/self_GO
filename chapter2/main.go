package main 
import "fmt"

func main () {
	str1 := "Hello, ";
	str2 := "Word!";
	number, text := add(3, 4, str1, str2);
	for i := 0; i < number; i++ {
		fmt.Println(text);
	}
}

func add(x, y int, first, last string) (int, string) {
	return x+y, first+last;
}	