package main

import "fmt"



func getSequence() func() int {
	i:=0
	return func() int { //此处即为闭包函数
	   i+=1
	  return i  
	}
 }
 
func main() {
	const (
		a = iota
		b 
		c 
	)

	fmt.Print(a,b,c)


}

