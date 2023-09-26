package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	// add a book to the library
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthorName(name string) []book {
	// return the list of books by the author
	return l[name]
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	jb := author{name: "James Baldwin"}

	lib.addBook(book{
		title:  "The Fire Next Time",
		author: jb,
	})

	lib.addBook(book{
		title:  "Go Tell It on the Mountain",
		author: jb,
	})

	// add 1 book to the library by a different author
	lib.addBook(book{
		title:  "Meditations",
		author: author{name: "Marcus Aurelius"},
	})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	books := lib.lookupByAuthorName(jb.name)

	// spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
	}
}
