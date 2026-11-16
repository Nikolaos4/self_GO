package main
import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 3

	fmt.Println(cap(intCh))     // 3
    fmt.Println(len(intCh))     // 1

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}