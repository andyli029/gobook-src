package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("hello")
	fmt.Printf("%x.\n", md5.Sum(data))
}
