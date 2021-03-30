package goker

import "testing"

func checkDuplicatedCards(cards []Card) *Card {
	check := make(map[Rank]map[Suit]bool)

	for _, card := range cards {
		rank := card.Rank()
		suit := card.Suit()

		if check[rank][suit] {
			return &card
		}

		check[rank][suit] = true
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
