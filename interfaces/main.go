package main

import "fmt"

//Interface declaration
type bot interface {
	getGreeting() string
	getVersion() float64
}

//structs
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	printBotVersion(eb)
	printBotVersion(sb)

}

//Print Method
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printBotVersion(b bot) {
	fmt.Println(b.getVersion())
}

//implementing the intrfaces methods for both the struts to qualify as bots
//Reciever funcs
func (eb englishBot) getGreeting() string {
	return "Hello There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!!"
}

func (eb englishBot) getVersion() float64 {
	return 5.6
}

func (sb spanishBot) getVersion() float64 {
	return 1.7
}
