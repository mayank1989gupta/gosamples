package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Some error while HTTP request.")
		os.Exit(1)
	}
	//Body --> io.ReadCloser --> Reader/Closer
	//creating a byte slice
	/*bs := make([]byte, 99999)
	//Read function reads the data into the byte slice provided
	//response body reads the data into the provided bs -> sliceis reference type hence it gets updated without pointer
	resp.Body.Read(bs)
	//fmt.Println(string(bs)) */
	fmt.Println("************ Using Helper **********************")
	//The same could be achieved by util/helper methods
	io.Copy(os.Stdout, resp.Body)

}
