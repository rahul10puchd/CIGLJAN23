package main

import (
	"fmt"
	"time"
)

// func main() {
// 	channel := make(chan string)
// 	go hello(channel)
// 	value := "main to anonymous"
// 	channel <- value //send value to channel
// 	//channel<-value //send value to channel
// 	fmt.Println("Exit main function")

// }
// func hello(channel chan string) {
// 	receive1 := <-channel //receive value from channel	\
// 	fmt.Println(receive1)
// 	time.Sleep(time.Second * 5)
// 	fmt.Println("In anonymous function")
// }
// func main() {
// 	channel := make(chan string)
// 	go send(channel)
// 	go receive(channel)
// 	channel <- "main to receive"
// 	fmt.Println(<-channel)
// 	fmt.Println("Exit main function")

// }

func hello(channel chan string) {
	receive1 := <-channel //receive value from channel	\
	fmt.Println(receive1)
	time.Sleep(time.Second * 5)
	fmt.Println("In anonymous function")
}
func main() {
	ch2 := make(chan string)
	channel1 := make(chan string, 1)
	//go send(sendingchannel)
	go receive(ch2)
	//sendingchannel <- "main to receive"
	channel1 <- "main to receive"
	//fmt.Println(<-channel1)
	fmt.Println("Exit main function")

}

func receive(channel <-chan string) {
	receive1, isTrue := <-channel //receive value from channel	\
	fmt.Println(receive1, isTrue)

	fmt.Println("In receive function")
}
func send(channel chan<- string) {

	//invalid operation: cannot receive from send-only channel channel (variable of type chan<- string)
	//receive1 := <-channel //receive value from channel	\
	channel <- "send to main"
	fmt.Println("In send function")
}
