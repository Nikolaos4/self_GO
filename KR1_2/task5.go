package main
import "fmt"

func main() {
	my_slice := []int{5, 8, 2, 10, 3, 7, 1, 9};
	new_slice := make([]int, len(my_slice))
	for i := 0; i < len(my_slice); i++ {
		if (my_slice[i] > 5) {
			new_slice = append(new_slice, my_slice[i]);
			fmt.Println(my_slice[i]);
		}
	}
}