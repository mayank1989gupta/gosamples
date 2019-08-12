package main

import "fmt"

//We can initialize a variable outside of a function, we just can't assign a value to it.
//For eg:
//var card string --> This is valid but we cannot assign a value

func main() {
	//defining variable in go --> All 3 are equivalent
	//var card string = "Ace of Cards"
	//var card = "Ace of Cards"
	// ":=" - Tells go compiler to figure about type and creation -->":=" is only for defining new variable
	//card := "Ace of Cards" //":=" only for creating new variable/tells go that we are creating a new variable
	card, card1 := newCard()
	fmt.Println(card + "," + card1)

	//creating slices --> where 'deck' is of type slice of []string -> deck.go file
	cards := newDeck()

	// deal method to return 2 decks -> deal &, playing cards
	hand, remainingCards := deal(cards, 2)
	hand.print()           //Calling the receiver function
	remainingCards.print() //Calling the receiver function

	//for loop --> i is index, card is the current object. range cards -> slice
	// we reinitialize i & card everytime as with for loop and every iteration value is thrown away
	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/

	//The above for loop id moved to the receiver function of cards
	//cards.print()

	var laptop laptopSize = 4.65
	fmt.Println("Size of Laptop: ", laptop.getSizeOfLaptop())

	//Type conversions in go
	fmt.Println(cards.toString())
	cards.saveToFile("deck")
	fmt.Println("*********Saved to file************")
	fmt.Println("*********Reading from file********")
	cardsFromFile := newDeckFromFile("deck")
	//cardsFromFile.print()

	fmt.Println("*********Shuffling the deck********")
	cardsFromFile.shuffle()
	cardsFromFile.print()
}

//function declaration -> func --> functionname --> return type
//Returning --> mulitple values
func newCard() (string, string) {
	return "5 of Diamonds", "6 of Spade"
}
