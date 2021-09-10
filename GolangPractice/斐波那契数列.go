package main

import "fmt"

func fibonacci(n int64) (result int64) {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var a int64
	fmt.Scanln(&a)
	fmt.Println(fibonacci(a))

}
