package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	//Reading from the given filename
	//Approach 1
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error while reading file", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
	fmt.Println("********************************************")
	//Approach 2
	//Using open func &, io.Copy with Reader/Writer
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error while opening the file.", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
