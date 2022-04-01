package main

import (
	"log"
	"os"
	"time"
)

const refreshAPITime = 5 * time.Minute

func main() {
	ch := make(chan bool)
	go getData(ch, refreshAPITime)
	for success := range ch {
		if success {
			runServer()
		} else {
			log.Print("Can't run server without requested data")
			os.Exit(0)
		}
	}
}
