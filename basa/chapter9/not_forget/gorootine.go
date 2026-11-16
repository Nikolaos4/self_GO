package main
import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	work := func(n int) {
		fmt.Printf("Status: %d\n", n)
		wg.Done()
	}
	for i := 0; i < 2; i++ {
		go work(i)
	}
	wg.Wait()
}