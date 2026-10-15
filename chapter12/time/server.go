package main
import (
	"fmt"
	"net"
)

var dict = map[string] string {
	"blue": "синий",
	"red": "красный",
	"yeelow": "желтый",
}

func main() {
	listen, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server is listening")
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go GetAnswer(conn)
	}
}


func GetAnswer(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if err !=  nil {
			fmt.Println(err)
			break
		}
		answer, ok := dict[string(input[:n])]
		if ok == false {
			answer = "not in dict"
		}
		conn.Write([]byte(answer))
	}	
}	