package main
import (
	"fmt"
	"os"
	"io"
)

func main() {
	text := []byte("Vlad is my best friend")
	file, err := os.Create("hello.txt")
	if (err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(text)
	fmt.Println("Done")

	//----------------------

	f, err := os.Open("hello.txt")
	if (err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} 
	defer f.Close()
	data := make([]byte, 64)
	for {
		n, err := f.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}