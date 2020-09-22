package main

import "fmt"
import "project_two/src/usecase"

func main() {

	deck := usecase.GenerateDeck()
	fmt.Println("Printing the full deck: \n", deck, "\n")

	shuffled := usecase.ShuffleDeck(deck)

	fmt.Println("Printing the shuffled deck: \n", shuffled, "\n")

	// dealing and printing 4 hands
	p1, p2, p3, p4 := usecase.DealHands(shuffled)
	fmt.Println("Player 1: ", p1)
	fmt.Println("Player 2: ", p2)
	fmt.Println("Player 3: ", p3)
	fmt.Println("Player 4: ", p4)
}
