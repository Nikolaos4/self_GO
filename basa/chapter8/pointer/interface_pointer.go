package main
import "fmt"

type write interface {
	read()
}

type File struct {
	text string
}

func (f *File) read(new_text string) {
	f.text = new_text
}

func main() {
	file := File{"Hello, World!"}
	p := &file
	fmt.Println(file)
	p.read("Hello, work!!")
	fmt.Println(file)
}