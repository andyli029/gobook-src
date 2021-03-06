package main

import (
	"fmt"
	"time"
	//"time"
)

func worker(die chan bool) {
	fmt.Println("Begin: This is Worker.")
	i := 0
	for {
		//fmt.Printf("for again: %d.\n", i)
		select {
		//case xx：
		//做事的分支
		case <-die:
			//time.Sleep(time.Second * 2)
			fmt.Printf("Done: This is Worker:%d.\n", i)
			die <- true
			//return
			i++
			break
		}
	}
}

func main() {
	die := make(chan bool)

	go worker(die)
	time.Sleep(time.Second * 2)
	die <- true
	<-die
	fmt.Println("Worker goroutine has been terminated")

	die <- true
	<-die
	fmt.Println("Worker goroutine has been terminated2.")
}
