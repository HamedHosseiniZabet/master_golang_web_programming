package main

import (
	"fmt"
)

func main()  {
	go helloWorld()
	//time.Sleep(time.Millisecond * 1)
}

func helloWorld(){
	fmt.Println("Hello Bitches ")
}


