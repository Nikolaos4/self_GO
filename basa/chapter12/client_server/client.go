package main
import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer conn.Close()
	
	for {
		var str string
		fmt.Println("Input a word:")
		_, err := fmt.Scanln(&str)
		if err != nil {
			fmt.Println("Uncorrected type of input")
			continue
		}
		if n, err := conn.Write([]byte(str)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Перевод: ")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil || n == 0 {
			fmt.Println(err)
			return
		}
		new_value := string(buff[:n])
		fmt.Println(new_value)
	}
}