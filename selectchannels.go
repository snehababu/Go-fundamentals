//select example
package main

import(
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	
	go func(){
		ch1 <- 42
	}()

	//channel will do?
	//once one of these channels become available and this is true for both sending & recieving, it will enter the specific case
	select{
	case val := <-ch1 :
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2 :
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("-------")

	//we create a channel of float64 and then we create a function that is going to sleep 100 millisecond and send the value of 5 to the channel
	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
		//if the goroutine finished in time
	case val := <- out:
		fmt.Printf("got %f\n", val)
		//else timeout
	case <- time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}