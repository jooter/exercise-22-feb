package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	go func() {
		for range time.Tick(time.Second) {
			procWithRecover()
		}
	}()

	select {}
}

func procWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recover proc with", r)
		}
	}()

	proc()
}

func proc() {
	log.Println("running in proc")
	panic("ok")
}
