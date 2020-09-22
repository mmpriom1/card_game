package model

import "project_two/src/enum"

type Game struct {
	Deck []Card
}

type Card struct {
	Suit   enum.SuitType
	Number enum.FaceType
}

// declaring the global variable

var (
	SUITS = []enum.SuitType{
		enum.Club, enum.Diamond, enum.Spade, enum.Heart,
	}

	NUMBERS = []enum.FaceType{
		enum.Ace, enum.Two, enum.Three, enum.Four, enum.Five, enum.Six, enum.Seven,
		enum.Eight, enum.Nine, enum.Ten, enum.Jack, enum.Queen, enum.King,
	}
)
