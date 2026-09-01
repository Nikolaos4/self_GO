package main
import "fmt"

func print(n int, ch chan bool) {
	fmt.Printf("Горутина %d завершила свое выполнение \n", n)
	ch <- true
}

func main() {
	ch := make(chan bool)
	count := 4
	for i := 0; i < count; i++ {
		go print(i, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("The end")
}