package main
import "fmt"

func main() {
	intch := make(chan int)
	strch := make(chan string)

	go func() {
		intch <- 5
	}()
	fmt.Println(<-intch)

	go help(2, strch)
	fmt.Println(<-strch)
}

func help(n int, ch chan string ) {
	ch <- "Hello, world"
}