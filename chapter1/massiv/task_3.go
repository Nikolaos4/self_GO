package main
import (
	"fmt"
	"math/rand/v2"
)

func print_massiv(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j :=  0; j < 3; j++ {
			fmt.Print(arr[i][j], " ");
		}
		fmt.Println();
	}
	fmt.Println();
} 

func main() {
	var arr [3][3]int;
	arr_up := arr;
	arr_down := arr;
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++  {
			num := rand.IntN(20);
			arr[i][j] = num;
			if (i > j) {
				arr_down[i][j] = num;
			} else {
				arr_up[i][j] = num;
			}
		}
	}

	print_massiv(arr);
	print_massiv(arr_up);
	print_massiv(arr_down);
}

