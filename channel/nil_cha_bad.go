package main

import "fmt"
import "time"

func main() {
	var c1, c2 chan int = make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		c2 <- 7
		close(c2)
	}()

	for {
		select {
		case x := <-c1:
			fmt.Println(x)
		case x := <-c2:
			fmt.Println(x)
		}
	}
	fmt.Println("over")
}
