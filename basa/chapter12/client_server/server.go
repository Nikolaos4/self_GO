package main
import (
	"fmt"
	"net"
) 

var dict = map[string]string {
	"red": "красный",
	"blue": "синий",
	"yellow": "желтый",
}

func main() {
	server, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}	

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if err != nil || n == 0 {
			fmt.Println("Reading is finished")
			break
		}
		value, ok := dict[string(input[:n])]
		if ok == false {
			fmt.Println("Value is not found in dict")
		}
		fmt.Println("Value = ", value)
		conn.Write([]byte(value))
	}
}