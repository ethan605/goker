package goker

import (
	"reflect"
	"testing"
)

func checkDuplicatedCards(cards []Card) *Card {
	check := make(map[string]bool)

	for _, card := range cards {
		if check[card.String()] {
			return &card
		}

		check[card.String()] = true
	}

	return nil
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if dealtCards := deck.DealtCards(); len(dealtCards) > 0 {
		t.Errorf("NewDeck() initiated with non-empty dealt cards %q", dealtCards)
	}

}

func TestDeckDeal(t *testing.T) {
	deck := NewDeck()

	tables := []struct {
		cardsNum        int
		dealtCards      int
		totalDealtCards int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{1, 1, 2},
		{5, 5, 7},
		{99, 45, 52}, // All cards dealt
		{99, 0, 52},  // All cards dealt
	}

	for _, table := range tables {
		dealtCards := deck.Deal(table.cardsNum)
		totalDealtCards := deck.DealtCards()

		if len(dealtCards) != table.dealtCards {
			t.Errorf(
				"deck.Deal(%d): expect %d cards dealt in turn, got %d",
				table.cardsNum, table.dealtCards, len(dealtCards),
			)
		}

		if len(totalDealtCards) != table.totalDealtCards {
			t.Errorf(
				"deck.Deal(%d): expect %d cards dealt in total, got %d",
				table.cardsNum, table.totalDealtCards, len(totalDealtCards),
			)
		}

		if duplicatedCard := checkDuplicatedCards(totalDealtCards); duplicatedCard != nil {
			t.Errorf("deck.Deal(%d): %q was dealt duplicately", table.cardsNum, *duplicatedCard)
		}
	}
}

func TestShuffledCards(t *testing.T) {
	orderedCards := []Card{}

	for _, rank := range allRanks {
		for _, suit := range allSuits {
			card, _ := NewCard(rank, suit)
			orderedCards = append(orderedCards, card)
		}
	}

	deck := NewDeck()
	deck.Deal(deckSize)
	shuffledCards := deck.DealtCards()

	tables := []struct{ from, to int }{
		{0, 3},
		{5, 10},
		{11, 21},
		{0, deckSize},
	}

	for _, table := range tables {
		orderedChunk := orderedCards[table.from:table.to]
		shuffledChunk := shuffledCards[table.from:table.to]

		if reflect.DeepEqual(orderedChunk, shuffledChunk) {
			t.Errorf("Deck wasn't shuffled well enough: %q (range %d %d)", shuffledChunk, table.from, table.to)
		}
	}
}
