package main
import "fmt"
import "sync"

var wg sync.WaitGroup

func get_num(ch <- chan int) {
	fmt.Println(<-ch)
	wg.Done()
}

func give_num(ch chan <- int) {
	ch <- 33
	wg.Done()
}

func main() {
	wg.Add(2)
	chanel := make(chan int)
	go get_num(chanel)
	chanel <- 22

	chanel2 := make(chan int)
	go give_num(chanel2)
	fmt.Println(<-chanel2)
}