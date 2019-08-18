package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}
	//creating a channel
	c := make(chan string)

	for _, link := range links {
		//creating go routines
		go checkStatus(link, c)
	}
	//whenever we are waiting for channel -> this is blocking channel &, Main Go Routine si put to sleep
	//fmt.Println(<-c)

	//To print channels --> recieving from channel 'c'
	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/

	//Repeating go-routines
	/*for {
		// With the below statement we are executing the func again with the value 'link'
		// which is passed by the above method call in the for loop
		// The above method call returns the link and the below fires the method with link = "<-c" fetched
		// channel
		go checkStatus(<-c, c)
	} */

	//The above for loop could be replaced with
	//Go identifies the type channel, here l is the link
	for l := range c {
		//Adding sleep/pause the current go routine
		//time.Sleep(5 * time.Second) //pauses for 5 sec -> This puts main go routine to sleep for 5s
		//go checkStatus(l, c)
		//Using function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l) //() is added to make sure the func is executed --> l is passed from outside func to func literal
	}
}

//Method to check the status of the given link
// put channel as parametere string has to be mentioned: c chan string
func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//sending the data into channel
		//c <- "Might be down"
		c <- link
		return
	}

	fmt.Println(link, "website is ok!")
	//c <- link + " -> website is ok"
	c <- link
}
