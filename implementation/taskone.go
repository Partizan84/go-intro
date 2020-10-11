package implementation

import (
    "fmt"
	"log"
	"strconv"
)

//Book Объявляем тип Book, который удовлетворяет интерфейсу fmt.Stringer.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

//Count Объявляем тип Count, который удовлетворяет интерфейсу fmt.Stringer.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

//WriteLog Объявляем функцию WriteLog(), которая берёт любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
    log.Println(s.String()," - сборника рассказов.")
}

//Books Инициализируем объекты Book и Count.
func Books() {
    // Инициализируем объект Book и передаём в WriteLog().
    book := Book{"Сказания Меекханского пограничья", "Роберт М. Вегнер"}
	WriteLog(book)
	
	// Инициализируем объект Count и передаём в WriteLog().
    count := Count(4)
    WriteLog(count)

}