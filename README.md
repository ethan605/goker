# goker

Poker machine, written in Go

**Table of Contents**
- Documentation
	- [Flows](./docs/FLOWS.md)
- [Components](#components)
	- [Cards](#cards)
	- [Decks](#decks)

## Components

### Cards

Cards are presented by public interface `Card` with read-only getters:

```go
type Card interface {
	Rank() Rank
	Suit() Suit
}
```

In which:

```go
type Rank int // from 2 to 10, Jack, Queen, King and Ace
type Suit string // "Club", "Diamond", "Heart" and "Spade"
```

The only way to access to a Card instance is via the [Decks](#decks)

### Decks

A deck is a collection of cards, with 2 representative information:
	- Cards remaining in stack
	- Cards being dealt

The `Deck` interface provides an accessible way to "deal" with cards:

```go
type Deck interface {
	// Get all dealt cards
	DealtCards() []Card

	// Deal a number of cards, return the drawn cards in this deal only.
	// To access all dealt cards, use `DealCards()`.
	// If there're not enough cards, deal all the remaining ones.
	Deal(int) []Card
}
```

Notes: only dealt cards are accessible, remaining cards are always enclosed.

Example for playing with decks:

```go
package main

import (
	"fmt"

	"github.com/ethan605/goker"
)

func main() {
	deck := goker.NewDeck()
	fmt.Println(deck) // goker.Deck<dealt: 0, remaining: 52>
	fmt.Println(deck.DealtCards())

	newCards := deck.Deal(1)
	fmt.Println(deck)                                  // goker.Deck<dealt: 1, remaining: 51>
	fmt.Println("New dealt cards:", newCards)          // 1 card dealt in this turn
	fmt.Println("All dealt cards:", deck.DealtCards()) // 1 card dealt in total

	newCards = deck.Deal(5)
	fmt.Println(deck)                                  // goker.Deck<dealt: 6, remaining: 46>
	fmt.Println("New dealt cards:", newCards)          // 5 cards dealt in this turn
	fmt.Println("All dealt cards:", deck.DealtCards()) // 6 cards dealt in total

	newCards = deck.Deal(99)
	fmt.Println(deck)                                  // goker.Deck<dealt: 52, remaining: 0>
	fmt.Println("New dealt cards:", newCards)          // 46 cards dealt in this turn
	fmt.Println("All dealt cards:", deck.DealtCards()) // 52 cards dealt in total

	newCards = deck.Deal(1)
	fmt.Println(deck)                                  // goker.Deck<dealt: 52, remaining: 0>
	fmt.Println("New dealt cards:", newCards)          // 0 cards dealt in this turn
	fmt.Println("All dealt cards:", deck.DealtCards()) // 52 cards dealt in total
}
```
