package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		//creating go routines
		go checkStatus(link)
	}
}

//Method to check the status of the given link
func checkStatus(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "website is ok")

}
