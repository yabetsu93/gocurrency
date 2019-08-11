package utils

import (
	"fmt"
	"time"
)

// MessageRelays message relays through channels
func MessageRelays(c1 chan string, c2 chan string, c3 chan string) (chan string, chan string, chan string) {
	go func() {
		// for the go routine to run continuously every 6 ms,
		// it has to be inside of the for loop and do the job
		// this will run and stop once it reach 0
		for i := 10; i >= 0; i-- {
			c1 <- "run every 9000 ms"
			time.Sleep(time.Millisecond * 9000)
			fmt.Printf("time left for iteration running every 9000 ms = %v", i)
		}
	}()

	go func() {
		// for the go routine to run continuously every 2 sec,
		// it has to be inside of the for loop and do the job
		// this will run and stop once it reach 0
		for i := 10; i >= 0; i-- {
			c2 <- "run every 90 seconds"
			time.Sleep(time.Second * 90)
			fmt.Printf("time left for iteration running every 90 seconds = %v", i)
		}
	}()

	go func() {
		// for the go routine to run continuously every 2 mins,
		// it has to be inside of the for loop and do the job
		// this will run and stop once it reach 0
		for i := 10; i > 0; i-- {
			c3 <- "run every minute 2 mins"
			time.Sleep(time.Minute * 2)
			fmt.Printf("time left for iteration running every 2 mins = %v", i)
		}
	}()

	return c1, c2, c3
}
