// Package cards provides a set of utilities for working with playing cards.

package cards

import (
	"math/rand"
)

// Type Suit represents a suit
type Suit byte

const (
	Clubs    Suit = 'C'
	Diamonds Suit = 'D'
	Hearts   Suit = 'H'
	Spades   Suit = 'S'

//	Suit(0) Suit = 0
)

func (s Suit) String() string {
	return string([]byte{byte(s)})
}

// Type Card represents a playing card
// An unknown card is represented by Card("")
type Card string

// Function New creates a new card from a given suit and value
func New(value uint8, suit Suit) Card {
	valuestr := map[uint8]string{
		1:  "A",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "J",
		12: "Q",
		13: "K",
	}[value]

	if valuestr == "" {
		valuestr = "0"
	}
	return Card(valuestr + suit.String())
}

// Method Suit returns the suit if there is any or Suit(0) otherwise
func (c Card) Suit() Suit {
	if len(c) < 2 {
		return Suit(0)
	}
	suit := Suit(c[len(c)-1])
	switch suit {
	case Clubs, Diamonds, Hearts, Spades:
		return suit
	default:
		return Suit(0)
	}
}

// Method Value returns the face value of the card, or 0 if there is none.
// Ace is represented by 1; Jack, Queen, and King are represented by 11, 12, and 13, respectively.
func (c Card) Value() uint8 {
	if len(c) < 2 {
		return 0
	}
	value := string(c[:len(c)-1])
	return map[string]uint8{
		"1":  1,
		"A":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
	}[value]
}

// Method String() returns the string representation of a card. Use c.String() instead of string(c) when possible.
func (c Card) String() string {
	if string(c) == "" {
		return "?"
	}
	return string(c)
}

// Type Deck represents a deck of playing cards
type Deck []Card

// Function FullDeck creates and returns a full deck of cards
func FullDeck() Deck {
	deck := make(Deck, 0, 52)
	for value := uint8(1); value <= 13; value++ {
		deck = append(deck, New(value, Clubs))
		deck = append(deck, New(value, Diamonds))
		deck = append(deck, New(value, Hearts))
		deck = append(deck, New(value, Spades))
	}
	return deck
}

// Method Shuffle shuffles a deck
func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
