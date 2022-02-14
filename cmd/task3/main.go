package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	ch := make(chan int)

	go func() {
		producer(ch)
	}()

	consumer(ch)
}

func producer(ch chan<- int) {
	d := 0
	for range time.Tick(time.Second) {
		d++
		log.Println("producer is about to send data:", d)
		ch <- d
	}
}

func consumer(ch <-chan int) {
	for d := range ch {
		log.Println("data has been received by consumer:", d)
	}
}
