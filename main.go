package main

import (
	"fmt"
	"net/http"
	"flag"
)

func main() {
	flag.Parse()
	urls := flag.Args()
	

	for _, url := range urls {
		status, err := checkURL(url)

		if err != nil {
			fmt.Printf("%s: %s\n", url, err)
		} else {
			fmt.Printf("%s: %v\n", url, status)
		}
	}
}

func checkURL(url string) (int, error){
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}
