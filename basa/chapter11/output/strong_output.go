package main
import (
	"fmt"
	"os"
	"io"
)

type person struct {
	name string
	age int
	weight float64
}

func writer(filename string) {
	persons := []person{
		{"Tom", 20, 66.5},
		{"KAte", 20, 50.5},
		{"Ann", 30, 50.0},
	}

	file, err := os.Create(filename)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	for _, p := range persons {
		fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
	}
}

func reader(filename string) {
	var name string
	var age int
	var weight float64

	file, err := os.Open(filename)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	for {
		_, err := fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
		if (err != nil) {
			if (err == io.EOF) {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} 
		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}
}

func main() {
	filename := "persons.txt"
	writer(filename)
	reader(filename)
}