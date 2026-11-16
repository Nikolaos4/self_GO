package main

import (
	"fmt"
)

func isPrime(n int) bool{
	flag := true;
	for i := 2; i < n/2; i++ {
		if (n % i == 0 ) {
			flag = false;
		}
	}
	if (flag) {
		return true;
	} else {
		return false;
	}
}

func main() {
	for i := 2; i < 101; i++ {
		if (isPrime(i)) {
			fmt.Println(i);
		}
	}
}