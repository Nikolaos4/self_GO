/*Метод представляет функцию, связанную с определенным типом. Методы определяются также как и обычные функции 
за тем исключением, что в определении метода также необходимо указать получателя или receiver. 
Получатель - это параметр того типа, к которому прикрепляется метод:

func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
    тело_метода
}
*/

package main
import "fmt"

type list []string


func (l list) print() {
	for _, value := range l {
		fmt.Print(value, " ");
	}
}

//

type person struct {
	name string
	age int
}

func (human person) print_person() {
	fmt.Println("Name of person: ", human.name);
	fmt.Println("Age of person: ", human.age);
}	

func (human person) characteristic(chr string) string{
	return human.name + " is a " + chr + " boy";
} 

func main() {
	var l list = list{"Tom", "Cat", "home"};
	l.print();
	//
	current_human := person{name: "Nick", age: 20};
	current_human.print_person();
	
	fmt.Println(current_human.characteristic("cool"));
}