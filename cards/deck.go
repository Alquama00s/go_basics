package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Diamond", "Hearts", "Club"}

	values := []string{"Ace", "two", "three", "four"}

	cards := deck{}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func fromString(s string) deck {

	return strings.Split(s, ",")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 4; i < len(d); i++ {

		randNumber := r.Intn(i)

		d[i], d[randNumber] = d[randNumber], d[i]

	}

}

func newDeckFromFile(filename string) deck {

	contentbyte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return fromString(string(contentbyte))
}
