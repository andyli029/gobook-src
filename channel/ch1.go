// http://tonybai.com/2014/09/29/a-channel-compendium-for-golang/
package main

import "fmt"
import "time"

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something…")

		time.Sleep(2 * time.Second)
		//c <- true
		close(c)
	}()
	<-c
	fmt.Println("Done!")
}
