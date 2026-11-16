package main
import (
	"fmt"
	"strconv"
)

func (book Book) GetInfo() string {
	return "Название: " + book.title + ", Автор: " + book.author + ", Год: " + strconv.Itoa(book.year);
}

func (book *Book) SetYear(new_year int) {
	book.year = new_year;
}

type Book struct {
	title string
	author string
	year int
}

func main() {
	my_book := Book {title: "Cat", author: "Tom", year: 3};
	data := my_book.GetInfo();
	fmt.Println(data);

	my_book.SetYear(5);
	new_data := my_book.GetInfo();
	fmt.Println(new_data);
}