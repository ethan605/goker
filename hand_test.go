package goker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandGetters(t *testing.T) {
	hand := &handStruct{}

	assert.NotEqual(t, hand.ID(), "", "unexpected Hand.ID() value")
	assert.Equal(t, len(hand.Cards()), 0, "unexpected Hand.Cards() value")
	assert.Equal(t, hand.CurrentBet(), 0, "unexpected Hand.CurrentBet() value")
}
