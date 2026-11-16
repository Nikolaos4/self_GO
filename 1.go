package main
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int = 0

func help(id int, m *sync.Mutex) {
	m.Lock()
	counter = 0
	for i := 0; i < 5; i++ {
		counter += i
		time.Sleep(10*time.Millisecond)
		fmt.Printf("Gorutine %d - %d:\n", id, counter)
	}

	m.Unlock()
	wg.Done()
}

func main() {
	var m sync.Mutex
	wg.Add(4)

	for i := 1; i <= 4; i++ {
		go help(i, &m)
	}
	wg.Wait()
}