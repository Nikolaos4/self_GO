package main
import "fmt"

type phone_number struct{}

func(ph phone_number) Writer(b []byte) (int, error){
	if len(b) == 0 {
		return 0, nil
	}
	for i := 0; i < len(b); i++ {
		if (b[i] >= '0' && b[i] <= '9') {
			fmt.Print(string(b[i]))
		}
	}
	fmt.Println()
	return len(b), nil
}

func main() {
	b1 := []byte("+1(234)567 9010")
	b2 := []byte("+2-345-678-12-35")

	p1 := phone_number{}
	p1.Writer(b1)
	p1.Writer(b2)
}