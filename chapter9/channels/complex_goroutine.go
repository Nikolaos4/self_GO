package main
import "fmt"

func sq(ch chan int) {
	num := <- ch
	ch <- num * num
}

func main() {
	ch := make(chan int)
	go sq(ch)
	ch <- 4
	fmt.Println(<-ch)
}