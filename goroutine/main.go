package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)

	// -----------------
	goroutine("sync")

	go goroutine("asnyc1")
	go goroutine("asnyc2")
	go goroutine("asnyc3")

	time.Sleep(time.Second * 3)
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	channel <- true
}

func goroutine(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text, "***", i)
	}
}
