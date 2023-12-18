/*
We are working on a program that allows users to store and retrieve information about
their favourite books. We have implemented a Book struct to store the information for the book.

You need to implement a function that will change the value of the Pages field for a given Book.

Also, make the required changes in the main function.
*/

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	book.Pages = pages
}

func main() {
	// Create 3 book structures
	book1 := Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180,
	}
	book2 := Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
		Pages:  281,
	}

	book3 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Pages:  279,
	}

	// Update the information for Books
	updatePages(&book1, 210)
	updatePages(&book2, 250)
	updatePages(&book3, 295)

	// Print all the struct objects
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)
}
