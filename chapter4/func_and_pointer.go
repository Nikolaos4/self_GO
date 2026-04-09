package main
import "fmt"

func Point(x *int) {
	*x = *x * *x;
}

func get_point(x int) *int {
	p := new(int);
	*p = x;
	return p;
}

func main() {
	x := 2;
	fmt.Println(x);
	Point(&x);
	fmt.Println(x);

	//указатели как результат функции
	p := get_point(3);
	fmt.Println(p);
	fmt.Println(*p);
}