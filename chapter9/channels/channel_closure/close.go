//если канал закрыт, то мы н можем с ним функионировать
package main
import "fmt"

func send_value(ch chan int) {
	ch <- 10
	close(ch)
}

func get_value(ch chan int) {
	value, ok := <- ch
	if (ok) {
		fmt.Println(value)
	} else {
		fmt.Println("Chanel is close")
	}
}

func main() {
	intch := make(chan int)
	go send_value(intch)
	get_value(intch)
	get_value(intch)
}