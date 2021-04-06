package goker

import (
	"fmt"
)

// Rank of card in a standard Poker deck.
// Available values: 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace
type Rank int

// Ranks enum
const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Suit of card in a standard Poker deck.
// Available values: Club, Diamond, Heart, Spade
type Suit string

// Suits enum
const (
	Club    Suit = "♣"
	Diamond Suit = "♦"
	Heart   Suit = "♥"
	Spade   Suit = "♠"
)

// Card represents information about a standard Poker card
type Card interface {
	fmt.Stringer

	// Standard ranks: 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace
	Rank() Rank

	// Standard suits: Club, Diamond, Heart, Spade
	Suit() Suit
}

func (card cardStruct) String() string {
	return fmt.Sprintf("goker.Card<%s%s>", namedRanks[card.rank-Two], card.suit)
}

func (card cardStruct) Rank() Rank {
	return card.rank
}

func (card cardStruct) Suit() Suit {
	return card.suit
}

/* Private stuffs */

var _ Card = (*cardStruct)(nil)

var (
	allRanks   = [...]Rank{2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace}
	namedRanks = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	allSuits   = [...]Suit{Club, Diamond, Heart, Spade}
)

type cardStruct struct {
	rank Rank
	suit Suit
}
