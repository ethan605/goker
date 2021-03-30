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
)

type _Card struct {
	rank Rank
	suit Suit
}

type Card interface {
	Rank() Rank
	Suit() Suit
}

var _ Card = &_Card{}

func (card _Card) Rank() Rank {
	return card.rank
}

func (card _Card) Suit() Suit {
	return card.suit
}

func (rank Rank) validate() error {
	switch {
	case rank >= Two && rank <= King:
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
		return _Card{}, err
	}

	if err := suit.validate(); err != nil {
		return _Card{}, err
	}

	return _Card{rank, suit}, nil
}
