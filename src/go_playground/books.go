package main

import "fmt"

type books struct {
	title, author, subject string
	bookID                 int
}

func printBook(book books) {

	fmt.Printf("Book Title : %s\n", book.title)
	fmt.Printf("Book Author :%s\n", book.author)
	fmt.Printf("Book Subject :%s\n", book.subject)
	fmt.Printf("Book Id :%d\n", book.bookID)
}
