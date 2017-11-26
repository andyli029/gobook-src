// http://tonybai.com/2014/09/29/a-channel-compendium-for-golang/
package main

import "fmt"

func main() {
	cb := make(chan bool)
	close(cb)
	x := <-cb
	fmt.Printf("%#v\n", x)

	x, ok := <-cb
	fmt.Printf("%#v %#v\n", x, ok)

	ci := make(chan int)
	close(ci)
	y := <-ci
	fmt.Printf("%#v\n", y)

	cb <- true
}
