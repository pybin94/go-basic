package main

import (
	"fmt"
	"time"
)

func main() {
	go count("oops")
	count("viin")
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is ", i)
		time.Sleep(time.Second)
	}
}
