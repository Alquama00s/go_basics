package main

import "fmt"

func main() {

	cards := newDeckFromFile("cardsfile")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	cards.print()
	cards.shuffle()
	cards.print()
	fmt.Println(cards.toString())
	// er := cards.saveToFile("cardsfile")

	// fmt.Println(er)
}
