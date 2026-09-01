package main
import "fmt"

func double(ch chan int) {
	num := <-ch
	ch<- num + num
}

func sq(ch chan int) {
	num := <-ch
	ch<- num*num
}

func cub(ch chan int) {
	num := <-ch
	ch<- num*num*num
}

func main() {
	ch_sq := make(chan int)
	ch_cub := make(chan int)
	ch_double := make(chan int)

	go double(ch_double)
	go sq(ch_sq)
	go cub(ch_cub)

	ch_sq<- 3 
	ch_cub<- 4
	ch_double<- 2
	for i := 0; i < 3; i++ {
		select {
		case sq_val := <-ch_sq:
			fmt.Println(sq_val)
		case d_val := <-ch_double:
			fmt.Println(d_val)
		case cub_val := <-ch_cub:
			fmt.Println(cub_val)
		}
	}
}