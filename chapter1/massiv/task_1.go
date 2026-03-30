package main
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var arr [6]int;
	var start, end = 0, 0;

	fmt.Print("Номер билета: ");
	for i := 0; i < 6; i++ {
		num := rand.IntN(10);
		arr[i] = num;
		fmt.Print(num, " ");
		if i < 3 {
			start += num;
		} else {
			end += num;
		}
	}

	fmt.Println("\nСумма первых 3-х: ", start);
	fmt.Println("Сумма последних 3-х: ", end);
	if (start == end) {
		fmt.Println("Билет счастливый");
	} else {
		fmt.Println("Билет несчастливый");
	}
}