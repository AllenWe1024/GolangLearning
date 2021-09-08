package main

import "fmt"

func main() {
	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}

	book1 := Books{"111", "222", "333", 4}
	//Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}

	fmt.Println(book1)
}
