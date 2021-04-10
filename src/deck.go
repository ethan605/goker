package goker

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck contains 2 stacks of Cards: remaining & dealt cards. Remaining cards are always enclosed.
// At initialization, there are no dealt cards. After each turn of Deal(),
// some of the remaining cards are disclosed and moved to the dealt stack.
type Deck interface {
	fmt.Stringer

	// Get all dealt cards
	DealtCards() []Card

	// Deal a number of cards, return the drawn cards in this deal only.
	// To access all dealt cards, use `DealCards()`.
	// If there're not enough cards, deal all the remaining ones.
	Deal(int) []Card
}

func (deck deckStruct) String() string {
	return fmt.Sprintf("goker.Deck<dealt: %d, remaining: %d>", len(deck.dealtCards), len(deck.remainingCards))
}

func (deck deckStruct) DealtCards() []Card {
	return deck.dealtCards
}

func (deck *deckStruct) Deal(cardsNum int) []Card {
	// Max number of available cards to draw is size of remaining cards
	numDrawnCards := len(deck.remainingCards)

	if cardsNum < numDrawnCards {
		numDrawnCards = cardsNum
	}

	drawnCards := make([]Card, numDrawnCards)

	drawnCards, deck.remainingCards = deck.remainingCards[:numDrawnCards], deck.remainingCards[numDrawnCards:]
	deck.dealtCards = append(deck.dealtCards, drawnCards...)

	return drawnCards
}

// NewDeck creates a new Deck with fully shuffled cards
func NewDeck() Deck {
	deck := assembleDeck()
	return &deckStruct{remainingCards: shuffleDeck(deck)}
}

/* Private stuffs */

const deckSize = len(allRanks) * len(allSuits)

type deckStruct struct {
	remainingCards []Card
	dealtCards     []Card
}

var _ Deck = (*deckStruct)(nil)

func shuffleCards(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func assembleDeck() []Card {
	deck := []Card{}
	ch := make(chan Card)

	// Loop through all ranks & suits in goroutines,
	// then append the cards to the deck randomly
	for _, rank := range allRanks {
		go func(rank Rank) {
			for _, suit := range allSuits {
				go func(suit Suit) {
					ch <- cardStruct{rank, suit}
				}(suit)
			}
		}(rank)
	}

	for i := 0; i < deckSize; i++ {
		deck = append(deck, <-ch)
	}

	return deck
}

func shuffleDeck(deck []Card) []Card {
	ch := make(chan []Card)

	// Shuffle the whole deck 3 times to make sure the cards order is fully random
	go func() {
		for i := 0; i < 3; i++ {
			shuffleCards(deck)
		}

		ch <- deck
	}()

	return <-ch
}
