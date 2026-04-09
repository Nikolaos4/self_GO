package main
import "fmt"

func main() {
	var x int = 4;
	var pointer *int = &x;
	fmt.Println("Адресс в памяти:", pointer);
	fmt.Println("Полученное значение: ", *pointer);

	//изменение значения:
	*pointer = 25;
	fmt.Println("Новое значение: ", *pointer);

	//указатель на указатель
	new_pointer := &pointer;
	fmt.Println(new_pointer, *new_pointer, **new_pointer)

	//массив указателей
	var a, b, c = 1, 2, 3;
	arr_point := [4]*int {&a, &b, &c}

	for _, v := range arr_point {
		fmt.Println(v, " ");
	}
}