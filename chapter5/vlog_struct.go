package main
import "fmt"

//вложенная структура
type person struct {
	name string
	age int
}

type purchase struct {
	product string
	price int
	person_data person
}

func main() {
	cat := purchase {
		product: "FatCat",
		price: 1000000,
		person_data: person {
			name: "Nick",
			age: 100,
		},
	}

	fmt.Println(cat);
	fmt.Println(cat.person_data.name);
}