package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{
			url:     url,
			err:     nil,
			latency: t,
		}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)

	urls := []string{
		"https://go.dev/",
		"https://go.dev/solutions/case-studies",
		"https://go.dev/doc/effective_go",
		"https://go.dev/security/",
	}

	for _, url := range urls {
		go get(url, results)
	}

	for range urls {
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s, %s\n", r.url, r.latency)
		}
	}
}
