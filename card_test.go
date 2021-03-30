package goker

import "testing"

func TestNewCard(t *testing.T) {
	cardZero := _Card{}

	tables := []struct {
		rank Rank
		suit Suit
		card _Card
		err  string
	}{
		{1, Heart, cardZero, "Invalid rank"},
		{15, Spade, cardZero, "Invalid rank"},
		{10, "Fake", cardZero, "Invalid suit"},
		{10, Heart, _Card{10, Heart}, ""},
		{10, Heart, _Card{10, Heart}, ""},
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
		{_Card{10, Heart}, 10, Heart},
		{_Card{3, Spade}, 3, Spade},
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
