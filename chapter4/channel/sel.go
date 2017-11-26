package main

import (
	"fmt"
	"time"
)

func worker(die chan bool, index int) {
	fmt.Println("Begin: This is Worker:", index)
	for {
		select {
		//case xx：
		//做事的分支
		case <-die:
			fmt.Println("Done: This is Worker:", index)
			return
		}
	}
}

func main() {
	die := make(chan bool)

	for i := 1; i <= 10; i++ {
		go worker(die, i)
	}

	time.Sleep(time.Second * 2)
	close(die)
	time.Sleep(time.Second * 1)
	//select {} //deadlock we expected
}
