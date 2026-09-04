//Замыкание (closure) представляет объект функции, который запоминает свое лексическое окружение даже в том случае, когда она выполняется вне своей области видимости.
package main
import "fmt"

func multiply(n int) func(int) int{
	return func (m int) int {return m * n}
}

func main() {
	result := multiply(5)(6)
	fmt.Println(result)
}