package main
import (
	"fmt"
	"os"
	"io"
)

func main() {
	text := "hello world! \n"
	text2 := []byte("Goodbye, my friend")
	file, err := os.Create("hello.txt")
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(text)
	file.Write(text2)

	fmt.Println("Done")

	//------------------------

	File_op, err := os.Open("hello.txt")
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	defer File_op.Close()
	data := make([]byte, 64)
	for {
		n, err := File_op.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}