package main
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("We can not open file. Error: ", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
}