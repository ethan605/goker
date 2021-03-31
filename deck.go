package goker

import (
	"math/rand"
	"time"
)

type Deck interface {
	// Get all dealt cards
	DealtCards() []Card

	// Deal a number of cards, return the drawn cards in this deal only.
	// To access all dealt cards, use `DealCards()`.
	// If there're not enough cards, deal all the remaining ones.
	Deal(int) []Card
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

func NewDeck() Deck {
	return &deckStruct{remainingCards: initDeck()}
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

func initDeck() []Card {
	deck := []Card{}
	cardsChannel := make(chan Card)

	// Loop through all ranks & suits in goroutines, cards will be appened to deck randomly
	for _, rank := range allRanks {
		go func(rank Rank) {
			for _, suit := range allSuits {
				go func(suit Suit) {
					cardsChannel <- cardStruct{rank, suit}
				}(suit)
			}
		}(rank)
	}

	for i := 0; i < deckSize; i++ {
		deck = append(deck, <-cardsChannel)
	}

	shuffleChannel := make(chan []Card)

	// Shuffle the whole deck 3 times to make sure the cards order is fully random
	go func() {
		for i := 0; i < 3; i++ {
			shuffleCards(deck)
		}

		shuffleChannel <- deck
	}()

	return <-shuffleChannel
}
