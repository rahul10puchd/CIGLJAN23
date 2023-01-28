package main
import(
	"fmt"
)
// func main(){
// 	channel:=make(chan string)
// 	go func(){
// 		value:="Hello World"
// 		channel<-value //send value to channel
// 	}()

// 	receive:=<-channel //receive value from channel	
// 	fmt.Println(receive)
// }
// func main(){
// 	channel:=make(chan string,1)
// 	go func(){
// 		receive:=<-channel //receive value from channel	\
// 		fmt.Println(receive)
// 		fmt.Println("In anonymous function")
// 	}()
// 	value:="main to anonymous"
// 	channel<-value //send value to channel
// 	channel<-value //send value to channel
// 	fmt.Println("Exit main function")
	
// }


// func main(){
// 	channel:=make(chan string,1)
// 	go func(){
// 		receive:=<-channel //receive value from channel	
// 		fmt.Println(receive)
// 		fmt.Println("In anonymous function")
// 	}()
// 	//value:="main to anonymous"
// 	// channel<-value //send value to channel
// 	fmt.Println("Exit main function")
	
// }


func mainnn(){
	channel:=make(chan string,1)
	go func(){
		receive1:=<-channel //receive value from channel	\
		fmt.Println(receive1)
		receive2:=<-channel //receive value from channel	\
		fmt.Println(receive2)
		fmt.Println("In anonymous function")
	}()
	value:="main to anonymous"
	channel<-value //send value to channel
	//channel<-value //send value to channel
	fmt.Println("Exit main function")
	
}