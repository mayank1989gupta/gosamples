package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type 'deck'
// which is slice of strings
type deck []string
type laptopSize float64

//receiver functions --> receives the elements of the types.
//Any variable of type deck will get access to this func
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//function to return new deck --> This is not a receiver function
func newDeck() deck {
	//cards := deck{}
	var cards deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//If we dont want to use the index -> we replace it with "_" -> which tells go that we will not use the variable
	for _, cardSuit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+cardSuit)
		}
	}

	return cards //return new deck of card
}

//Deal function --> to return the deal and playing cards.
func deal(d deck, handSize int) (deck, deck) {
	set1 := d[:handSize]
	set2 := d[handSize:]

	return set1, set2
}

//Method to return the String separated by "," from the given slice using 'strings' package
func (d deck) toString() string {
	return strings.Join(d, ",")
}

//Method to save to a given file -> returns error if any
//As writeFile expects slice of byte, we convert deck type to String and use type conversion to
//convert from deck ([] string) to string
func (d deck) saveToFile(fileName string) error {
	//Type conversion from String to byte slice "[]byte" &, '0666' means anyone can read/write the file
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//Method to read slice of bytes from given file and convert to String and split on ","
func newDeckFromFile(filename string) deck {
	//variables for byte slice &, error object
	bs, err := ioutil.ReadFile(filename)
	//In case of error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//If no error execute the below
	cardsString := string(bs)
	result := strings.Split(cardsString, ",")

	return deck(result)
}

//Method to shuffle the deck &, generate random number
func (d deck) shuffle() {
	//This is required as the ran.Intn will always genearte the same randon number
	// so we update the source by passing the time.Now().UnixNano()
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//where i is the index
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//swapping in go --> a, b = b, a
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (ls laptopSize) getSizeOfLaptop() laptopSize {
	return ls
}
