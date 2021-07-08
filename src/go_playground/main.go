package main

import "fmt"
import "go_basics/logics"

func main() {

	/***************** Book Entity Block *****************************/

	var book1 books
	var book2 books

	book1.title = "Trigonometry"
	book1.author = "R.D. Sharma"
	book1.subject = "Maths"
	book1.bookID = 451202

	book2.title = "Go Programming"
	book2.author = "Go Author"
	book2.subject = "GO"
	book2.bookID = 789456

	printBook(book1)
	printBook(book2)

	/*****************************************************************/

	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	shapes := []shape{r, c}

	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(shape.perimeter())
	}

	measure(r)
	measure(c)

	logics.Palindrome(999)
	logics.FindLongConsecutive()
}
