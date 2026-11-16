package main
import "fmt"

type Vehicle interface{
	move() //метод, который будет переопределен в зависимости от конкретного типа
}

//структура автомобиль
type Car struct {}

//структура самолет
type Aircraft struct {}

func (c Car) move() {
	fmt.Println("Машина едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	var tesla Vehicle = Car{} //переменная интерфейса
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
}