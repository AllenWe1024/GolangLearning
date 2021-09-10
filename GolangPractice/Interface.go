package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type N1Phone struct {
}

func (n1 N1Phone) call() {
	fmt.Println("n1's call")
}

type N2Phone struct {
}

func (n2 N2Phone) call() {
	fmt.Println("n2's call")
}

func main() {
	var phone1 Phone

	phone1 = new(N1Phone)
	phone1.call()

	phone1 = new(N2Phone)
	phone1.call()
}
