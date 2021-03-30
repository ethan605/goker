package goker

type Suit string

const (
	Heart   Suit = "Heart"
	Spade        = "Spade"
	Club         = "Club"
	Diamond      = "Diamond"
)

type Rank int

const (
	One Rank = iota
)

type Card struct {
	suit Suit
}

func NewCard() Card {
	card := Card{suit: Heart}

	return card
}
