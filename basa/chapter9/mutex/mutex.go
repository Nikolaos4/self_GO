package main
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int = 0


func change(n int, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for i := 0; i < n; i++ {
		counter += i
		fmt.Println(n, ":", counter)
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	var mutex sync.Mutex
	wg.Add(4)

	for i := 1; i <= 4; i++ {
		go change(i, &mutex)
	} 
	wg.Wait()
	fmt.Println("The end")
}