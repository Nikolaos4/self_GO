package main
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprint(file, "Сегодня ")
	fmt.Fprintln(file, "хорошая погода")
}