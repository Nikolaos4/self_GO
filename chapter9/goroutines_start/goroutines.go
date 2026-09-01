package main
import "fmt"

func sum(n int) {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	fmt.Println(n, "-", result)
}

func main() {
	for i := 0; i < 7; i++ {
		go sum(i)
	}
	fmt.Scanln()
	fmt.Println("The end")

}