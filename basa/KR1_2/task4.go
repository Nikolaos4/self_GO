package main

import (
	"fmt"
	"math/rand/v2"
)

func summ(arr [10] int) (string, int) {
	sum := 0;
	for i := 0; i < len(arr); i++ {
		sum += arr[i];
	}
	return "Суима =", sum;
}

func main() {
	var arr [10] int;
	for i := 0; i < 10; i++ {
		arr[i] = rand.IntN(51);
		fmt.Print(arr[i], " ");
	}

	message, summa := summ(arr);
	fmt.Printf("\n%s %d. Среднее арифметическое: %d", message, summa, summa/10);
}