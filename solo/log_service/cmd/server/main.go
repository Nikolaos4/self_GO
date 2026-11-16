package main
import (
	"fmt"
	"net"	
	"bufio"
	"strings"
	"tcp_logs/internal/protocol"
)


func GetRequest(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		str := scanner.Text()
		str = strings.TrimSpace(str)
		arr_from_str := strings.Fields(str)
		result := protocol.ProtocolCheck(arr_from_str)
		if result == "-1" {
			conn.Close()
			return
		}
		fmt.Fprintf(conn, "%s\n", result)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8181")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server is listenning")
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go GetRequest(conn)
	}
}

//нужно понять, как я буду хранить данные о количестве записей для каждого из логов