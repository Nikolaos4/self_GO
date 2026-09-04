package main
import (
	"fmt"
	"os"
)

type person struct{
	name string
	age int32
	weight float64
}

func main() {
	Tom := person{
		name: "Tom",
		age: 30,
		weight: 55.5,
	}
	file, err := os.Create("person.txt")
	if err != nil {
		fmt.Println(file)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "%-10s %-10d %-10.3f\n", Tom.name, Tom.age, Tom.weight)
}