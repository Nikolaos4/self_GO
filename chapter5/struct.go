//структуру можно определять как внутри main так и перед ним
package main
import "fmt"

type FatCat struct {
	name string
	weight int
}

func get_pointer(cat *FatCat) {
	cat.name = "Mircha";
}

func main() {
	var Mira FatCat = FatCat{"Mira", 8};
	fmt.Println(Mira);

	Ley := FatCat{name: "Ley", weight: 7};
	fmt.Println(Ley.name, Ley.weight);

	//анонимная струткура

	person := struct {
		name string
		age int
	} {
		name: "Tom",
		age: 23,
	}
	fmt.Println(person.name, person.age);

	//указатели

	type weather struct {
		string 
		description string
		int
	}

	today := weather {"Хорошая", "Пасмурно", 20};
	var pointer *weather = &today;
	fmt.Println(pointer.description);
	pointer.description = "Отлично";
	fmt.Println((*pointer).description);

	//функция new
	yesterday := new(weather);
	fmt.Println(yesterday);

	//передача структуры в функцию
	get_pointer(&Mira);
	fmt.Println(Mira);
}