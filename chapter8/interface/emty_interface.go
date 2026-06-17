package main
import "fmt"

type Emty interface{}

func print(object Emty) {
	fmt.Println(object)
}

type person struct {
	name string
}

type acount struct {
	id int
}

func main() {
	Tom := person{"Tom"}
	tom_ac := acount{125}

	//печать
	print(Tom)
	print(tom_ac)
}