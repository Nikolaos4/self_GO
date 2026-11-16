package main
import (
	"fmt"
)

func main() {
	intch := make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		intch <- 5
	} ()
	fmt.Println(<-intch)
	fmt.Println("The end")
}