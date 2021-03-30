package goker

import "testing"

func TestNewCard(t *testing.T) {
	cardZero := card{}

	tables := []struct {
		rank Rank
		suit Suit
		card Card
		err  string
	}{
		{1, "Heart", cardZero, "Invalid rank"},
		{15, "Spade", cardZero, "Invalid rank"},
		{10, "Fake", cardZero, "Invalid suit"},
		{7, "Heart", card{Seven, Heart}, ""},
		{11, "Club", card{Jack, Club}, ""},
		{13, "Diamond", card{King, Diamond}, ""},
		{14, "Spade", card{Ace, Spade}, ""},
	}

	for _, table := range tables {
		received, err := NewCard(table.rank, table.suit)

		switch {
		case err != nil && err.Error() != table.err:
			t.Errorf("NewCard() throws %q, want %q", err, table.err)

		case received != table.card:
			t.Errorf("NewCard() returns %q, want %q", received, table.card)
		}
	}
}

func TestCardGetters(t *testing.T) {
	tables := []struct {
		card Card
		rank Rank
		suit Suit
	}{
		{card{Seven, Heart}, 7, "Heart"},
		{card{Three, Spade}, 3, "Spade"},
	}

	for _, table := range tables {
		rank := table.card.Rank()
		suit := table.card.Suit()

		switch {
		case rank != table.rank:
			t.Errorf("Card.Rank() returns %q, want %q", rank, table.rank)

		case suit != table.suit:
			t.Errorf("Card.Rank() returns %q, want %q", suit, table.suit)
		}
	}
}
