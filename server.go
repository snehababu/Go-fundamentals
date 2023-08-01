/*
Services, or servers, usually store their process ID in a file. When you want to terminate the service, you need to read this process ID from the file and then terminate the service. You need to write a function called killServer. It should read the process ID from the file, convert it from a string to an integer, and finally, print out killing the server with process ID. You need to think about error handling and what to do in each case of an error. And then, in the main function, you're going to call the killServer with server.pid, and if there is an error, print out the error to standard error.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	//Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)

	//after killing the file we need to remove the server
	if err := os.Remove(pidFile); err != nil {
		//We can fo on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}
	return nil
}

func main(){
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}