package main
import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	text string
}

func (file *File) read() {
	fmt.Println(file.text)
}

func (file *File) write(message string) {
	file.text = message
	fmt.Println("Запсись в файл строки", message)
}

func main() {
	myFile := &File{}
	myFile.write("Hello, word")
	myFile.read() 
	myFile.write("By!")
	myFile.read()
}