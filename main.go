package main

func main() {

	cards := newDeck()
	// cards := newDeck() //deck{"Ace of spades", newCard()}

	// // cards.print()

	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()

	// // cards.print()

	// // greeting := "Hi there"
	// // fmt.Println([]byte(greeting))

	// // cards.print()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards2 := newDeckFromFile("my_cards")
	// cards2.print()

	cards.shuffle()
	cards.print()
}
