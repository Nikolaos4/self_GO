package main
import (
	"fmt"
	"os"
	"io"
)

type data string

func(d data) Read(b []byte) (int, error) {
	count := 0
	for i := 0; i < len(d); i++ {
		if (d[i] >= '0' && d[i] <= '9') {
			b[count] = d[i]
			count++
		}
	}
	return count, io.EOF
}

func Write(filename string) {
	d := data("Привет, Гари \nКак дела")
	file, err := os.Create(filename)
	if (err != nil) {
		fmt.Println(err)
		return 
	} 
	defer file.Close()
	fmt.Fprintln(file, d)
}

func Read(filename string) {
	var hi, how string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fscanln(file, &hi)
	fmt.Fscanln(file, &how)
	fmt.Println(hi, how, '\n', "Done")
}

func main() {
	name := "file.txt"
	Write(name)
	Read(name)
}