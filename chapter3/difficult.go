package main
import (
	"fmt"
	"math/rand/v2"
)

func add(slice [][]int) int {
	sum := 0;
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			sum += slice[i][j];
		}
	}
	return sum;
}

func mult_of_add(slice [][] int, operation func(slice [][]int) int) int {
	mult := 1;
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] % 2 == 0 {
				mult *= operation(slice);
			} else {
				mult = 1;
			}
		}
	}
	return mult;
}

func main() {
	num1 := rand.IntN(3)+2;
	num2 := rand.IntN(3)+2;
	
	slice := make([][]int, num1);
	for i := 0; i < num1; i++ {
		slice[i] = make([]int, num2);
		for j := 0; j < num2; j++ {
			number := rand.IntN(20);
			slice[i][j] = number;
			fmt.Print(number, " ");
		}
		fmt.Println();
	}

	var f1 func(slice [][]int) int = add;
	fmt.Println("\nСумма массива: ", f1(slice));

	f2 := mult_of_add;
	fmt.Print("\nПроизведение строк = ", f2(slice, add));
}