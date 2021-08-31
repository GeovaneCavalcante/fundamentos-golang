package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return ch
}

func main() {
	t1 := titulo("https://www.google.com")
	t2 := titulo("https://www.google.com")
	fmt.Printf("%s, %s", <-t1, <-t2)
	fmt.Printf("%s, %s", <-t1, <-t2)
}
