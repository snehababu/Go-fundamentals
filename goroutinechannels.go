package main

import (
	"fmt"
	"time"
)

/*
func main() {
	ch := make(chan int)
	ch <- 353 //send

	val := <-ch //recieve
	fmt.Printf("got %d\n", val)
}
//it creates a panic : we get all goroutines are asleep and deadlock
//when we are using a send operation, we are blocked until somebody is going to read from the channel
//go has a mechanism to deal with deadlocks and panic
//fix below
*/

func main() {
	ch := make(chan int)

//introducing a goroutine that will send a value to the channel. This goroutine will run concurrently to the main one which will recieve from the other channel and print it out 
	go func() {
		//send number of the channel
		ch <- 353
	}()

	val := <-ch //recieve
	fmt.Printf("got %d\n", val)

	fmt.Println("------")

	//send multiple values to the channel
	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i //send
			time.Sleep(time.Second)
		}
	}()

		for i :=0; i < count; i++ {
			val := <- ch //reieving value from the channel
			fmt.Printf("recieved %d\n", val)
		}

	fmt.Println("------")

	//close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch{
		fmt.Printf("received %d\n", i)
	}

	fmt.Println("------")	
	//buffered channel
	//this channel has buffer of 1, which means we can send one value without blocking
	//second value sending in is going to block
	ch2 := make(chan int, 1)
	ch2 <- 19
	val2 := <-ch2
	fmt.Println(val2)

}
