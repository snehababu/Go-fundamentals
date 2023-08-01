//get content type of sites
package main

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return 
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string) {
	for _,url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	//to know when the goroutines are done
	var wg sync.WaitGroup
	for _,url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main(){
	urls := []string {
		"https://golang.com",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	//serial
	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	//concurrent run faster than serial
	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))
}