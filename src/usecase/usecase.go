package usecase

import (
	"math/rand"
	"project_two/src/model"
	"time"
)
import "project_two/src/enum"

// NewCard will create a new card with suite and face value
func NewCard(suite enum.SuitType, value enum.FaceType) model.Card {
	return model.Card{
		Suit:   suite,
		Number: value,
	}
}

// GenerateDeck will take will create a full set of card (usually 52 cards) and will return it as a Game
func GenerateDeck() model.Game {
	game := model.Game{}
	for _, suite := range model.SUITS {
		for _, numbers := range model.NUMBERS {
			game.Deck = append(game.Deck, NewCard(suite, numbers))
		}
	}

	return game
}

// ShuffleDeck will take a game as input and will shuffle the card
func ShuffleDeck(d model.Game) []model.Card {
	cards := d.Deck

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

// DealHands will deal the cards between 4 hands
func DealHands(cards []model.Card) ([]model.Card, []model.Card, []model.Card, []model.Card) {

	p1 := make([]model.Card, 0)
	p2 := make([]model.Card, 0)
	p3 := make([]model.Card, 0)
	p4 := make([]model.Card, 0)

	for i := 0; i < len(cards); i += 4 {
		p1 = append(p1, cards[i])
		p2 = append(p2, cards[i+1])
		p3 = append(p3, cards[i+2])
		p4 = append(p4, cards[i+3])
	}

	return p1, p2, p3, p4
}
