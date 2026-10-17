package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
	Err    error
	Time int
}

func main() {
	flag.Parse()
	urls := flag.Args()

	var wg sync.WaitGroup

	ch := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			
			start := time.Now()

			status, err := checkURL(u)

			duration := time.Since(start)
			ch <- Result{URL: u, Status: status, Err: err, Time: int(duration.Milliseconds())}
		}(url)
	}

	wg.Wait()
	close(ch)

	for r := range ch {
		if r.Err != nil {
			fmt.Printf("%s: %s %v\n", r.URL, r.Err, r.Time)
		} else {
			fmt.Printf("%s: %v %v\n", r.URL, r.Status, r.Time)
		}
	}
}

func checkURL(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}
