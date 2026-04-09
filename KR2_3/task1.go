package main
import "fmt"

func main() {
	var name string;
	var age int;
	var height float64;

	fmt.Scanln(&name);
	fmt.Scanln(&age);
	fmt.Scanln(&height);

	fmt.Printf("%s, %d лет, рост %f", name, age, height);
}