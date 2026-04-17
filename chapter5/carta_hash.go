//карта - это та же хэш-таблица, которая предсталяет из себя структуру данных, представляющую из себя пары ключ-значение
//map[K]V - объявление карты, где K - тип ключа(тип должен поддерживать операцию сравнения ==), V - тип значения

package main
import (
	"fmt"
	"reflect"
)

func main() {
	var people = map[int]string{
		1: "Ann",
		2: "Bob",
		3: "Dan",
	}
	fmt.Println(people[1]);
	people[1] = "Alice";
	fmt.Println(people[1]);

	//проверка наличия значения по определенному ключу
	if i, ok := people[2]; ok {
		fmt.Println(i);
	}
	
	//перебор
	for key, value := range people {
		fmt.Println("Key: ", key, ", value: ", value);
	}

	//создание карты с помощью make
	var new_people = make(map[int]string);
	//добавение 
	new_people[1] = "Alice";
	//удаление
	delete(new_people, 1);

	//сравнение
	fmt.Println(reflect.DeepEqual(people, new_people));
}