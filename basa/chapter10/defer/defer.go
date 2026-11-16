package main
import "fmt"
import "os"

func open() {
	file, err := os.Open("./teoria.txt")
	defer file.Close()
	fmt.Println("file: ", file)
	fmt.Println("err: ", err)
}

func main() {
	open()
}