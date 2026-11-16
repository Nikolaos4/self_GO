package main
import "fmt"

func swap(a, b *int) {
	*b = (*a + *b);
	*a = *b - *a;
	*b = *b - *a;
}

func main() {
	var a, b = 2, 4;
	swap(&a, &b);
	fmt.Println(a, b);
}