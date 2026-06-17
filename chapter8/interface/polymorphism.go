package main
import "fmt"

type Vehicle interface {
	move()
}

type Car struct { model string}

type Aircraft struct { model string}

func (c Car) move() {
	fmt.Println(c.model, "едет")
}

func (a Aircraft) move() {
	fmt.Println(a.model, "летит")
}

func main() {
	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	boing := Aircraft{"Boing"}

	vehicles := [...] Vehicle{tesla, volvo, boing}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}