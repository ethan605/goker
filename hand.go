package goker

import (
	"fmt"
)

// Hand holds a pair of cards and some amount of bet in each turn
type Hand interface {
	fmt.Stringer

	// Get the hand's ID
	ID() string

	// Get the card pairs
	Cards() pair

	// Get the bet amount in current turn
	CurrentBet() int

	// End the current game, reset cards and bet
	endGame()

	// Receive a new card
	receiveCard(Card) error
}

func (hand handStruct) String() string {
	return fmt.Sprintf("goker.Hand")
}

func (hand handStruct) ID() string {
	return ""
}

func (hand handStruct) CurrentBet() int {
	return 0
}

func (hand handStruct) Cards() pair {
	return pair{}
}

/* Private stuffs */

var _ Hand = (*handStruct)(nil)

type pair [2]Card

type handStruct struct {
}

func (hand *handStruct) endGame() {
}

func (hand *handStruct) receiveCard(card Card) error {
	return nil
}
