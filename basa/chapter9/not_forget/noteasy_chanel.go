package main
import "fmt"

func main() {
	ch := make(chan int)
	go square(ch)
	ch <- 5
	fmt.Println(<-ch)

	fmt.Println(<-createchanel())
}

func square(ch chan int) {
	num := <-ch
	ch <- num*num
}

func createchanel() chan string {
	new_ch := make(chan string)
	go func() {
		new_ch <- "hello, world"
	}()
	return new_ch
}