package main

import (
	"fmt"
	"time"
)

// func count(thing string, c chan string) {

// 	for i := 0; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(1 * time.Second)
// 	}
// 	close(c)
// }
// func main() {
// 	c := make(chan string)
// 	go count("anish", c)
// 	for msg := range c {
// 		//reciver of the channel
// 		fmt.Println(msg)
// 	}
// }

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
