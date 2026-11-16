package main
import "fmt"

func main() {
	f_arr := []func(){}
	for i := 0; i < 5; i++ {
		j := i
		f_arr = append(f_arr, func() {
			fmt.Println(j)
		})
	}

	for _, val := range f_arr {
		val()
	}
}