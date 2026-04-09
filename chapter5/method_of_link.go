package main
import "fmt"

type person struct {
	name string
	age int
}

func (human *person) update_age(new_age int) {
	human.age = new_age;
}

func main() {
	Tom := person{"Tom", 20};
	fmt.Println("Age before: ", Tom.age);
	//var pointer *person = &Tom;
	//pointer.update_age(44);
	Tom.update_age(33)
	fmt.Println("Age after: ", Tom.age);
	//так проще!!!
}