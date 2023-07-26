//Writing a function that return Content-type header
package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedIn.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //Make sure we close the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		//Return error if the Content-Type header is not found
		return "", fmt.Errorf("can't find Content-Type header")
	}
	return ctype, nil
}