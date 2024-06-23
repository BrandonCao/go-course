package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, website := range websites {
		go checkLink(website, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", link, err)
		channel <- "Might be down I think"
		return
	}

	if resp.StatusCode == http.StatusOK {
		channel <- "it is up"

		fmt.Printf("%s is up and running\n", link)
	} else {
		channel <- "is down"
		fmt.Printf("%s is down with status code %d\n", link, resp.StatusCode)
	}
}
