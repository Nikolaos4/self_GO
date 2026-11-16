package main
import "fmt"

func main() {
	chanel := make(chan int, 3)
	chanel <- 3
	chanel <- 4
	close(chanel)

	for i := 0; i < cap(chanel); i++ {
		if val, opened := <- chanel; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Chanel closed")
		}
	}
}