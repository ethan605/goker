package goker

import "errors"

type Suit string

const (
	Heart   Suit = "Heart"
	Spade        = "Spade"
	Club         = "Club"
	Diamond      = "Diamond"
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

type cardStruct struct {
	rank Rank
	suit Suit
}

type Card interface {
	Rank() Rank
	Suit() Suit
}

var _ Card = &cardStruct{}

func (card cardStruct) Rank() Rank {
	return card.rank
}

func (card cardStruct) Suit() Suit {
	return card.suit
}

func (rank Rank) validate() error {
	switch {
	case rank >= Two && rank <= Ace:
		return nil
	}

	return errors.New("Invalid rank")
}

func (suit Suit) validate() error {
	switch suit {
	case Heart, Spade, Club, Diamond:
		return nil
	}

	return errors.New("Invalid suit")
}

func NewCard(rank Rank, suit Suit) (Card, error) {
	if err := rank.validate(); err != nil {
		return cardStruct{}, err
	}

	if err := suit.validate(); err != nil {
		return cardStruct{}, err
	}

	return cardStruct{rank, suit}, nil
}
