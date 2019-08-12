package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//Print Method
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//Reciever funcs
func (eb englishBot) getGreeting() string {
	return "Hello There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!!"
}
