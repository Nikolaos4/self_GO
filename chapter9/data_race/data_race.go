package main
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var value int = 22

func update() {
	fmt.Println("Value now:", value)
	if (value == 22) {
		time.Sleep(1*time.Second)
		fmt.Printf("Update value: %d \n", value)
		value++
	}
	fmt.Printf("Value after: %d \n", value)
	wg.Done()
}

func main() {
	wg.Add(2)

	go update()
	go update()

	wg.Wait()
}