package main

import (
	"fmt"
	"time"
)

func main() {
	helloChannel := make(chan string)
	goodByeChannel := make(chan string)
	quitChannel := make(chan bool)
	go receiveMessage(helloChannel, goodByeChannel, quitChannel)
	go sendMessage(helloChannel, "hello from Main")
	time.Sleep(time.Second * 5)
	go sendMessage(goodByeChannel, "good bye")
	<-quitChannel
}

func sendMessage(channel chan<- string, fromMain string) {
	channel <- fromMain
}
func receiveMessage(helloChannel, goodByeChannel <-chan string, quitChannel chan<- bool) {
	for {
		select {
		case msg := <-helloChannel:
			println(msg)
		case msg := <-goodByeChannel:
			println(msg)
		case <-time.After(time.Second * 6):
			fmt.Println("Nothing received in 6 seconds, Exiting Now")
			quitChannel <- true
			return
		}
	}
}
