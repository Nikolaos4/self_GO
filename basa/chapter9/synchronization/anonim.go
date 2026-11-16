package main
import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //создали счетчик выполняющихся горутин
	work := func(id int) {
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d закончила свое выполнение \n", id)
		wg.Done() //завершаем выполнение горутин
	}

	go work(2)
	go work(3)

	wg.Wait()

	fmt.Println("Горутины завершили свое выполнение")
}