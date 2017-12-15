package main

import "fmt"
import "time"

func worker(start chan bool, index int) {
	fmt.Println("start  Worker:", index)
	time.Sleep(1 * time.Second)
	<-start

	fmt.Println("stop Worker:", index)
}

func main() {
	start := make(chan bool)
	for i := 1; i <= 10; i++ {
		go worker(start, i)
	}
	time.Sleep(2 * time.Second)
	close(start)
	time.Sleep(3 * time.Second)
	//select {} //deadlock we expected
}
