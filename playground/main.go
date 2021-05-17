package main

import "fmt"

func main() {
	sages := []string{"Gandi", "MLK", "Buddah", "Jesus", "Mohammad"}
	for index, element := range sages {
		fmt.Println(index, element)
	}
}
