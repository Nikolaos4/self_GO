package main
import "fmt"

func factorial(n int, ch chan int) {
	defer close(ch) // закрывает канал, когда функция заканчивается
	result := 1
	for i := 1; i < n; i++ {
		result *= i
		ch<-result
	}
}

func main() {
	ch := make(chan int)
	go factorial(7, ch)
	for num := range ch { //пока канал не закрыт, можно работать с данными
		fmt.Println(num)
	}
}