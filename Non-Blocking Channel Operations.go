package main 

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Pringln("receved message",msg)
	default:
		fmt.Println("no message receive")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Pringln("sent message",msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg :=<-messages:
		fmt.Println("received message",msg)
	case sig := <-signals:
		fmt.Println("received signal",sig)
	default:
		fmt.Println("no activity")
	}
}