package main

import "github.com/NameSurname/Assignment1/Library"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	library := Library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add book")
		fmt.Println("2. Exit")
		fmt.Println("Choose and option: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter Book ID: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Book Title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter Book Author: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			library.AddBook(Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case 2:
			fmt.Println("Exiting... Bye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
