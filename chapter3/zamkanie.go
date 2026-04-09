//замыкание - функция, которая объединяет две вещи: функцию и окружение, в котором функция бла создана
package main 
import "fmt"

func zamk1() func(int) int {
	n := 5;
	return func(x int) int {return x + n}
}

func zamk2() func(){
	var n = 1;
	return func() {
		n += 1;
		fmt.Println(n)
	}
}

func main() {
	f := zamk1();
	fmt.Println(f(3));

	f2 := zamk2();
	f2();
	f2();
}	