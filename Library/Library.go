package Library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[string]Book
}

func NewLibrary() *Library {
	return &Library{make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	if _, exists := l.books[book.ID]; exists {
		fmt.Println("Book with such ID already exists.")
		return
	}
	l.books[book.ID] = book
	fmt.Println("Book was added.")
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.books[id]; exists {
		if book.IsBorrowed {
			fmt.Println("Book is borrowed.")
			return
		}
		book.IsBorrowed = true
		l.books[id] = book
		fmt.Println("You have borrowed the book.")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.books[id]; exists {
		if !book.IsBorrowed {
			fmt.Printf("Book is not currently borrowed.")
			return
		}
		book.IsBorrowed = false
		l.books[id] = book
		fmt.Println("You have returned the book")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	for _, book := range l.books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
