package main
import "fmt"

func zam() func() {
	n := 0
	return func() {
		n += 1
		fmt.Println(n)
	}
}

func main() {
	result := zam()
	result()
	result()
}