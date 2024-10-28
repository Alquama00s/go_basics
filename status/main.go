package main

import (
	"fmt"
	"net/http"
	"syscall"
)

func main() {
	v := 2 >> 1
	fmt.Print(v)
	fmt.Println("starting...")
	fmt.Println(syscall.Getpid())
	lnk := []string{
		"http://google.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range lnk {
		go checkLink(link, c)
	}

	for l := range c {

		go func(link string) {
			// time.Sleep(2000 * time.Millisecond)
			checkLink(link, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " down!")
		c <- link
	}
	fmt.Println(link + " up!")
	c <- link
}
