package main

func main() {

	cards := newDeck() //deck{"Ace of spades", newCard()}

	//cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
