package main
import "fmt"

type transport interface{
    move()
}

type plain struct{}

type car struct{}

func(p plain) move() {
    fmt.Println("fly")
}

func(c car) move() {
    fmt.Println("drive")
}

func main() {
    var p plain = plain{}
    var c car = car{}
    arr := [...]transport{p, c}
    for _, trans := range arr {
        trans.move()
    }
}