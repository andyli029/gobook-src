package main

import "fmt"

func newUniqueIDService() <-chan string {
	id := make(chan string)
	go func() {
		var counter int64 = 0
		for {
			id <- fmt.Sprintf("%d", counter)
			counter += 1
		}
	}()
	return id
}
func main() {
	id := newUniqueIDService()
	for i := 0; i < 11; i++ {
		fmt.Println(<-id)
	}
}
