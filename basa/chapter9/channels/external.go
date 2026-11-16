package main
import "fmt"

func main() {
	intch := make(chan int)
	go receive_numbers(5, intch)
	fmt.Println(<-intch)
}

func receive_numbers(n int, ch chan int) {
	ch <- n * n
}