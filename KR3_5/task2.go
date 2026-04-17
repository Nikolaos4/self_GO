package main
import "fmt"

func new_counter() func() {
	n := 0;
	return func() {
		n++;
		fmt.Println(n);
	}
}

func main() {
	c1 := new_counter();
	c1();
	c1();
	c1();
}