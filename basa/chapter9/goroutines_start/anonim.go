package main
import (
	"fmt"
	"time"
)

func main() {
	for i:= 0; i < 7; i++ {
		go func(n int) {
			result := 0
			for j := 1; j <= n; j++ {
				result += j
			}
			fmt.Println(n, "-", result)
		}(i)
	}
	time.Sleep(200*time.Millisecond)
	fmt.Println("The End")
}