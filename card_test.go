package goker

import (
	"fmt"
	"testing"
)

func TestNewCard(t *testing.T) {
	cardZero := cardStruct{}

	tables := []struct {
		rank Rank
		suit Suit
		card Card
		err  string
	}{
		{1, "Heart", cardZero, "Invalid rank"},
		{15, "Spade", cardZero, "Invalid rank"},
		{10, "Fake", cardZero, "Invalid suit"},
		{7, "Heart", cardStruct{Seven, Heart}, ""},
		{11, "Club", cardStruct{Jack, Club}, ""},
		{13, "Diamond", cardStruct{King, Diamond}, ""},
		{14, "Spade", cardStruct{Ace, Spade}, ""},
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
		str  string
	}{
		{cardStruct{7, Heart}, 7, "Heart", "7 of Heart"},
		{cardStruct{3, Spade}, 3, "Spade", "3 of Spade"},
		{cardStruct{Jack, Club}, 11, "Club", "Jack of Club"},
		{cardStruct{Queen, Diamond}, 12, "Diamond", "Queen of Diamond"},
		{cardStruct{King, Heart}, 13, "Heart", "King of Heart"},
	}

	for _, table := range tables {
		rank := table.card.Rank()
		suit := table.card.Suit()
		str := fmt.Sprint(table.card)

		switch {
		case rank != table.rank:
			t.Errorf("Card.Rank() returns %q, want %q", rank, table.rank)

		case suit != table.suit:
			t.Errorf("Card.Rank() returns %q, want %q", suit, table.suit)

		case str != table.str:
			t.Errorf("Card.String() returns %q, want %q", str, table.str)
		}
	}
}
