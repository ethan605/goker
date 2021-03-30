package goker

type deckStruct struct {
	dealtCards []Card
}

type Deck interface {
	DealtCards() []Card
	Deal(int) []Card
}

var _ Deck = (*deckStruct)(nil)

func (deck deckStruct) DealtCards() []Card {
	return deck.dealtCards
}

func (deck *deckStruct) Deal(cardsNum int) []Card {
	return deck.dealtCards
}

func NewDeck() Deck {
	return &deckStruct{}
}
