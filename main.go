package main

import "fmt"

func main() {
	c:=make(chan int,15)
	for i:=1;i<=10;i++ {
		c<-i
	}
	for i:=1;i<=10;i++ {
		fmt.Println(<-c)
	}
}

