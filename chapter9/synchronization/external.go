package main
import (
	"fmt"
	"time"
	"sync"
)

func executable_goroutines(id int, wg *sync.WaitGroup) {
	fmt.Printf("Горутина %d начала свое выполнение \n", id)
	time.Sleep(2*time.Second)
	fmt.Printf("Горутина %d закончила свое выполнение \n", id)
	wg.Done()
}	

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go executable_goroutines(2, &wg)
	go executable_goroutines(3, &wg)

	wg.Wait()
	fmt.Println("Горутины завершили свое выполнение")
}