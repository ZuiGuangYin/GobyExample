/*Channels are the pipes that connect concurrent goroutines.
*You an send values into another goroutine,Channels are a powerflul primitive that underly mach of Go's functionality.
*/

package main 

import "fmt"

func main() {
	messages := make(chan  string)

	go func() {messages <- "ping"}()

	msg := <-messages
	fmt.Println(msg)
}

