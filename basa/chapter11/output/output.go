package main
import (
	"fmt"
	"os"
)

func write(text string) {
	name := "Tom"
	age := 24

	file, err := os.Create(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func read(text string) {
	file, err := os.Open(text)
	var name string
	var age int
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}

func main() {
	text := "person.txt"
	write(text)
}