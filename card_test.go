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
		{cardStruct{7, Heart}, 7, "Heart", "goker.Card<7 of Heart>"},
		{cardStruct{3, Spade}, 3, "Spade", "goker.Card<3 of Spade>"},
		{cardStruct{Jack, Club}, 11, "Club", "goker.Card<Jack of Club>"},
		{cardStruct{Queen, Diamond}, 12, "Diamond", "goker.Card<Queen of Diamond>"},
		{cardStruct{King, Heart}, 13, "Heart", "goker.Card<King of Heart>"},
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
