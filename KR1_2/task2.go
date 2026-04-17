package main

import "fmt"

func main() {
	var num int;
	fmt.Scan(&num);
	
	if (num > 0) {
		fmt.Println("положительное");
	} else if (num < 0) {
		fmt.Println("отрицательное");
	} else {
		fmt.Println("ноль");
	}

}