// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	firstName string
	lastName  string
}

// define a book type with a title and author
type book struct {
	title      string
	authorName string
}

// define a library type to track a list of books
var library map[string][]book

// define addBook to add a book to the library
func addBook(title, authorname string) {
	library[authorname] = append(library[authorname], book{title: title, authorName: authorname})
}

// define a lookupByAuthor function to find books by an author's name
func lookupByAuthor(name string) []book {
	someBook, _ := library[name]
	return someBook
}

func main() {
	// create a new library
	library = make(map[string][]book)

	// add 2 books to the library by the same author
	addBook("Midsommer Night Dream", "Shakespeare")
	addBook("Romeo and Juliet", "Shakespeare")

	// add 1 book to the library by a different author
	addBook("A Little Princess", "Frances Hodgeson")

	// dump the library with spew
	//fmt.Printf("%#v\n", library)
	spew.Dump(library)

	books := lookupByAuthor("Shakespeare")
	// lookup books by known author in the library
	spew.Dump(books)

	// print out the first book's title and its author's name
	fmt.Printf("\nFirst book: %s by %s\n", books[0].title, books[0].authorName)

}
