//get content types of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string {
		"https://golang.com",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	//create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	//wait using channel
	for range urls { //run number of URLs times
		out := <- ch //block the channel until we get a value in the channel
		fmt.Println(out)
	}
}