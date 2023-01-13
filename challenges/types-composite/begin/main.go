// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
//
type author struct{
	name string
}

// define a book type with a title and author
//
type book struct{
	title string
	author author
}

// define a library type to track a list of books
//
type library map[string][]book

// define addBook to add a book to the library
//
func (l library)addBook(b book){
   l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
//
func (l library)lookupByAuthor(name string)[]book{
	return l[name]
}

func main() {
	// create a new library
	//
  lib := library{} 
	// add 2 books to the library by the same author
	//
	jb := author{name: "Kemal Yilmaz"}
	bk1 := book{
		title: "Book1",
		author: jb,
	}
	bk2 := book{
		title: "Book2",
		author: jb,
	}
	lib.addBook(bk1)
	lib.addBook(bk2)
	// add 1 book to the library by a different author
	//
	jb2 := author{name: "Mark Twain"}
	bk3 := book{
		title: "The Adventures of Tom Sawyer",
		author: jb2,
	}
	lib.addBook(bk3)
	// dump the library with spew
	//
	spew.Dump(lib)
	// lookup books by known author in the library
	//
  books :=lib.lookupByAuthor("Kemal Yilmaz")
	spew.Dump(books)
	// print out the first book's title and its author's name
	//
	if len(books)>0{
		fmt.Println(books[0].title + " by " + books[0].author.name)
	}
}
