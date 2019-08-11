package main

import (
	"fmt"

	utils "./goroutines/utils"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go utils.MessageRelays(c1, c2, c3)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg3 := <-c3:
			fmt.Println(msg3)
		}
	}
}
