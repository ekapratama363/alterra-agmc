package lib

import (
	"day-2/models"
	"fmt"
)

var books = make(map[string]models.Books)

func GetAllBook() map[string]models.Books {
	fmt.Println("books", books)
	return books
}

func CreateBook(book *models.Books) map[string]models.Books {
	books[book.Id] = *book
	return books
}
