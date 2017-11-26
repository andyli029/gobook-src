package main

import "fmt"
import "time"

func main() {
	var c1, c2 chan int = make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 7
		close(c2)
	}()

	for {
		select {
		case x, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println(x)
			}
		}

		if c1 == nil && c2 == nil {
			break
		}
	}
	fmt.Println("over")
}
