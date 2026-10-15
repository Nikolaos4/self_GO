package main
import (
	"fmt"
	"net"
	"bufio"
	"strings"
	//"sync"
)

var dict = map[string]string {
	"red": "красный",
	"blue": "синий",
	"yellow": "желтый",
}

//-------------------------------

type ComandDo interface {
	Execute(args []string) string
}

type Get struct{}
type Add struct{}
type List struct{}
type Quit struct{}		

func(Get) Execute(words []string) string{
	if len(words) > 1 {
		return "Error"
	}
	translate, ok := dict[words[0]]
	if ok == false {
		translate = "NotFound"
	}
	return translate
}

func(Add) Execute(words []string) string{
	if len(words) != 2 {
		return "Error! The collor is not create"
	}
	dict[words[0]] = words[1]
	return "Successful addition"
}

func(List) Execute(words []string) string{
	arr_from_dict := []string{}
	for key, _ := range dict {
		arr_from_dict = append(arr_from_dict, key)
	}
	answer_str := strings.Join(arr_from_dict, ", ")
	return answer_str
}

func(Quit) Execute(words []string) string{
	return ""
}

var comand_dict = map[string] ComandDo{
	"GET": Get{},
	"ADD": Add{},
	"LIST": List{},
	"QUIT": Quit{},
}

//--------------------------------

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
			continue
		}
		go getAnswer(conn)
	}
}

func getAnswer(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		str := scanner.Text()
		str = strings.TrimSpace(str)
		arr_words := strings.Fields(str)
		req, ok := comand_dict[arr_words[0]]
		if ok == false {
			fmt.Fprintf(conn, "Comand not found")
			return
		}
		arr_words = arr_words[1:]
		answer := req.Execute(arr_words)
		if answer == "" {
			break
		}
		fmt.Fprintf(conn, "%s\n", answer)
	}
}