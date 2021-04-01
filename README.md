# goker

Poker machine, written in Go

**Table of Contents**
1. [Components](#components)

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
