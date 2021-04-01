package goker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	dealtCards := deck.DealtCards()
	desc := fmt.Sprint(deck)

	assert.Equal(t, len(dealtCards), 0, fmt.Sprintf("NewDeck() should init with non-empty dealt cards, got %q", dealtCards))
	assert.Equal(t, desc, "goker.Deck<dealt: 0, remaining: 52>", "unexpected Deck.String() value")
}

func TestDeckDeal(t *testing.T) {
	deck := NewDeck()

	tables := []struct {
		cardsNum        int
		dealtCards      int
		totalDealtCards int
		desc            string
	}{
		{0, 0, 0, "goker.Deck<dealt: 0, remaining: 52>"},
		{1, 1, 1, "goker.Deck<dealt: 1, remaining: 51>"},
		{1, 1, 2, "goker.Deck<dealt: 2, remaining: 50>"},
		{5, 5, 7, "goker.Deck<dealt: 7, remaining: 45>"},
		{99, 45, 52, "goker.Deck<dealt: 52, remaining: 0>"}, // All cards dealt
		{99, 0, 52, "goker.Deck<dealt: 52, remaining: 0>"},  // No more cards dealt
	}

	for _, table := range tables {
		dealtCards := deck.Deal(table.cardsNum)
		totalDealtCards := deck.DealtCards()
		desc := fmt.Sprint(deck)
		duplicatedCard := checkDuplicatedCards(totalDealtCards)
		errMessage := fmt.Sprintf("unexpected Deck.Deal(%d) value", table.cardsNum)

		assert.Equal(t, len(dealtCards), table.dealtCards, errMessage)
		assert.Equal(t, len(totalDealtCards), table.totalDealtCards, errMessage)
		assert.Equal(t, desc, table.desc, errMessage)
		assert.Nil(t, duplicatedCard, errMessage)
	}
}

func TestShuffledCards(t *testing.T) {
	orderedCards := []Card{}

	for _, rank := range allRanks {
		for _, suit := range allSuits {
			orderedCards = append(orderedCards, cardStruct{rank, suit})
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
		errMessage := fmt.Sprintf("Deck wasn't shuffled well enough: %q (range %d %d)", shuffledChunk, table.from, table.to)

		assert.Equal(t, reflect.DeepEqual(orderedChunk, shuffledChunk), false, errMessage)
	}
}
