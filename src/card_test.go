package goker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardGetters(t *testing.T) {
	tables := []struct {
		card Card
		rank Rank
		suit Suit
		desc string
	}{
		{cardStruct{7, Heart}, 7, "♥", "goker.Card<7♥>"},
		{cardStruct{3, Spade}, 3, "♠", "goker.Card<3♠>"},
		{cardStruct{Jack, Club}, 11, "♣", "goker.Card<J♣>"},
		{cardStruct{Queen, Diamond}, 12, "♦", "goker.Card<Q♦>"},
		{cardStruct{King, Heart}, 13, "♥", "goker.Card<K♥>"},
	}

	for _, table := range tables {
		rank := table.card.Rank()
		suit := table.card.Suit()
		desc := fmt.Sprint(table.card)

		assert.Equal(t, rank, table.rank, "unexpected Card.Rank() value")
		assert.Equal(t, suit, table.suit, "unexpected Card.Suit() value")
		assert.Equal(t, desc, table.desc, "unexpected Card.String() value")
	}
}
