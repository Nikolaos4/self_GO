package main
import "fmt"

type Read interface{
	read()
}

type Writer interface{
	write(string)
}

type All interface{
	Read()
	witer(string)
}

type File struct{
	message string
}

func( )

func talk_about(file *All) {

}

func main() {
	file := File{"Hello, word!"}
	talk_about(file)
}