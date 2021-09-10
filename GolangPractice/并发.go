package main

import "fmt"

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", s)
	}
}

func main() {
	go say("Hello")
	say("World")
	say("!")
}
