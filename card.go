package goker

import (
	"errors"
	"fmt"
)

type Rank int

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

type Suit string

const (
	Club    Suit = "Club"
	Diamond Suit = "Diamond"
	Heart   Suit = "Heart"
	Spade   Suit = "Spade"
)

type Card interface {
	fmt.Stringer

	// Available ranks: 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace
	Rank() Rank

	// Available suits: Club, Diamond, Heart, Spade
	Suit() Suit
}

func (card cardStruct) Rank() Rank {
	return card.rank
}

func (card cardStruct) Suit() Suit {
	return card.suit
}

func (card cardStruct) String() string {
	return fmt.Sprintf("%s of %s", namedRanks[card.rank-Two], card.suit)
}

// Raise an error if either rank or suit is invalid
func NewCard(rank Rank, suit Suit) (Card, error) {
	if err := rank.validate(); err != nil {
		return cardStruct{}, err
	}

	if err := suit.validate(); err != nil {
		return cardStruct{}, err
	}

	return cardStruct{rank, suit}, nil
}

/* Private stuffs */

var _ Card = (*cardStruct)(nil)

var (
	allRanks   = [...]Rank{2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace}
	namedRanks = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	allSuits   = [...]Suit{Club, Diamond, Heart, Spade}
)

type cardStruct struct {
	rank Rank
	suit Suit
}

func (rank Rank) validate() error {
	if rank >= Two && rank <= Ace {
		return nil
	}

	return errors.New("Invalid rank")
}

func (suit Suit) validate() error {
	switch suit {
	case Heart, Spade, Club, Diamond:
		return nil

	default:
		return errors.New("Invalid suit")
	}
}
